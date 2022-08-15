package autowire

import (
	"fmt"
	"reflect"
	"sync"

	"ioc-go/autowire/util"
	"ioc-go/logger"

	perros "github.com/pkg/errors"
)

// 包装 Autowire，添加一些impl方法（增强）
type WrapperAutowire interface {
	Autowire

	// 无参数实现
	ImplWithoutParam(sdID string, withProxy bool) (interface{}, error)
	// 有参数实现
	ImplWithParam(sdID string, param interface{}, withProxy bool) (interface{}, error)

	implWithField(info *FieldInfo) (interface{}, error)
}

// 获取某个一个 wrapper autowire，传入Autowire(key) 和 WrapperAutowireMap
func getWrappedAutowire(autowire Autowire, allAutowires map[string]WrapperAutowire) WrapperAutowire {
	return &WrapperAutowireImpl{
		Autowire:           autowire,
		singletonImpledMap: map[string]interface{}{},
		allAutowires:       allAutowires,
	}
}

// wrapper autowire的实现
type WrapperAutowireImpl struct {
	Autowire
	singletonImpledMap     map[string]interface{}     // 单例Map(sdID-类型指针)，注：可以判断proxy是否开启，如果开启就会把类型指针做一层代理
	singletonImpledMapLock sync.RWMutex               // 读写锁（单例实现Map需要使用lock）
	allAutowires           map[string]WrapperAutowire // WrapperAutowireMap
}

// 获取无参数带结构实现
func (w *WrapperAutowireImpl) ImplWithoutParam(sdID string, withProxy bool) (interface{}, error) {
	param, err := w.ParseParam(sdID, nil)
	if err != nil {
		if w.Autowire.IsSingleton() {
			// FIXME: 忽略解析参数错误，因为带有空参数但单例也尝试从配置文件中查找属性
			logger.Blue("[Wrapper Autowire] Parse para from config file with sdid %s failed, error: %s, continue with nil param.", sdID, err)
			return w.ImplWithParam(sdID, param, withProxy)
		} else {
			return nil, err
		}
	}
	return w.ImplWithParam(sdID, param, withProxy)
}

// 获取带参数带结构实现
func (w *WrapperAutowireImpl) ImplWithParam(sdID string, param interface{}, withProxy bool) (interface{}, error) {
	// (1) check 单例
	w.singletonImpledMapLock.RLock()
	if singletionImpledPtr, ok := w.singletonImpledMap[sdID]; w.Autowire.IsSingleton() && ok {
		w.singletonImpledMapLock.RUnlock()
		return singletionImpledPtr, nil
	}
	w.singletonImpledMapLock.RUnlock()

	// (2) factory
	impledPtr, err := w.Autowire.Factory(sdID)
	if err != nil {
		return nil, err
	}
	if w.Autowire.InjectPosition() == AfterFactoryCalled { // 第一个注入factory后
		if err := w.inject(impledPtr, sdID); err != nil {
			return nil, err
		}
	}

	// (3) 构造函数属性（impledPtr 是一个接口）
	impledPtr, err = w.Autowire.Construct(sdID, impledPtr, param)
	if err != nil {
		return nil, err
	}
	if w.Autowire.InjectPosition() == AfterConstructorCalled { // 第二个注入construction 后
		if err := w.inject(impledPtr, sdID); err != nil {
			return nil, err
		}
	}

	// (4) 尝试包装代理
	if withProxy {
		// 如果字段是接口，尝试注入字段包装指针
		impledPtr = GetProxyFunction()(impledPtr)
	}

	// (5) 记录单例指针（save map）
	if w.Autowire.IsSingleton() {
		w.singletonImpledMapLock.Lock()
		w.singletonImpledMap[sdID] = impledPtr
		w.singletonImpledMapLock.Unlock()
	}
	return impledPtr, nil
}

// implWithField 用于从字段创建参数并调用 ImplWithParam
func (w *WrapperAutowireImpl) implWithField(info *FieldInfo) (interface{}, error) {
	sdID, err := w.ParseSDID(info)
	if err != nil {
		logger.Red("[Wrapper Autowire] Parse sdid for field %+v failed, error is %s", info, err)
		return nil, err
	}
	sd := GetStructDescriptor(sdID)
	// FIXME: 重大bug，对于全部实现 autowire，sd 可能为 nil
	// 类型类型是接口 && 禁用代理 才能返回 true
	implWithProxy := info.FieldReflectValue.Kind() == reflect.Interface && !sd.DisableProxy
	if err != nil {
		return nil, err
	}
	param, err := w.ParseParam(sdID, info)
	if err != nil {
		if w.Autowire.IsSingleton() { // 如果是单例
			// FIXME: 忽略解析参数错误，因为带有空参数的单例也尝试从配置文件中查找属性
			logger.Red("[Wrapper Autowire] Parse param from config file with sdid %s failed, error: %s, continue with nil param.", sdID, err)
			return w.ImplWithParam(sdID, param, implWithProxy)
		} else {
			return nil, err
		}
	}
	return w.ImplWithParam(sdID, param, implWithProxy)
}

// 注入做标记和 monkey 注入
func (w *WrapperAutowireImpl) inject(impledPtr interface{}, sdID string) error {
	sd := w.Autowire.GetAllStructDescriptors()[sdID]

	// (1) reflect反射
	valueOf := reflect.ValueOf(impledPtr)
	if valueOf.Kind() != reflect.Interface && valueOf.Kind() != reflect.Ptr { // 不是接口也不是指针直接返回
		return nil
	}
	valueOfElem := valueOf.Elem() // Elem() 返回值或者指针指向的值
	typeOf := valueOfElem.Type()
	if typeOf.Kind() != reflect.Struct { // 不是结构体，直接返回
		return nil
	}

	// 处理 struct
	// (2) 标记注入
	numField := valueOfElem.NumField()
	for i := 0; i < numField; i++ {
		field := typeOf.Field(i)
		var subImpledPtr interface{}
		var subService reflect.Value
		tagKey := ""
		tagValue := ""
		for _, aw := range w.allAutowires {
			if val, ok := field.Tag.Lookup(aw.TagKey()); ok {
				// check field
				subService = valueOfElem.Field(i)
				tagKey = aw.TagKey()
				tagValue = val
				if !(subService.IsValid() && subService.CanSet()) { //IsValid表示是否是零值（初始值）；CanSet表示是否可以修改
					errMsg := fmt.Sprintf("Failed to autowire struct %s's impl %s service, It's field type %s with tag '%s:\"%s\"', please check if the field name is exported",
						sd.ID(), util.GetStructName(impledPtr), field.Type.Name(), tagKey, tagValue)
					logger.Red("[Autowire Wrapper] Inject field failed with error: %s", errMsg)
					return perros.New(errMsg)
				}

				fieldType := buildFiledTypeFullName(field.Type)
				fieldInfo := &FieldInfo{
					FieldName:         field.Name,
					FieldType:         fieldType,
					TagKey:            aw.TagKey(),
					TagValue:          val,
					FieldReflectType:  field.Type,
					FieldReflectValue: subService,
				}
				// 从字段信息创建参数（field ——> param）
				var err error

				subImpledPtr, err = aw.implWithField(fieldInfo)
				if err != nil {
					return err
				}
				break // 只支持一个标签（注解）
			}
		}
		if tagKey == "" && tagValue == "" {
			continue
		}
		// FIXME: 设置属性（牛，先把subService已经设置了，才去给里面设置地址空间）。注意：这样会不会出现问题，还没有赋值地址空间，使用方调用了（NPE问题，Spring使用了二级缓存）
		subService.Set(reflect.ValueOf(subImpledPtr))
	}
	return nil
}

// TODO 可以将字段向下解析到 autowireImpl，但不是 autowire 的核心
func buildFiledTypeFullName(fieldType reflect.Type) string {
	// TODO 查找不支持的类型和日志警告，比如 "struct" 字段
	if util.IsPointerField(fieldType) || util.IsSliceField(fieldType) {
		return fieldType.Elem().PkgPath() + "." + fieldType.Elem().Name()
	}
	// 接口属性
	return fieldType.PkgPath() + "." + fieldType.Name()
}
