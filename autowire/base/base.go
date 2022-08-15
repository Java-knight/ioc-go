package base

import (
	perros "github.com/pkg/errors"
	"ioc-go/autowire"
	"ioc-go/logger"
)

// Autowire 自动装配接口
type FacadeAutowire interface {
	// SDMap
	GetAllStructDescriptors() map[string]*autowire.StructDescriptor
	TagKey() string
}

// 返回一个新的 autowire Base
func New(facadeAutowire FacadeAutowire, sp autowire.SDIDParser, pl autowire.ParamLoader) AutowireBase {
	return AutowireBase{
		facadeAutowire: facadeAutowire,
		sdIDParser:     sp,
		paramLoader:    pl,
	}
}

// autowire base info
type AutowireBase struct {
	facadeAutowire FacadeAutowire
	sdIDParser     autowire.SDIDParser
	paramLoader    autowire.ParamLoader
}

// 从SDMap[sdID] 找到 sd，返回 sd.Factory
func (this *AutowireBase) Factory(sdID string) (interface{}, error) {
	allStructDescriptor := this.facadeAutowire.GetAllStructDescriptors()
	if allStructDescriptor == nil {
		return nil, perros.New("struct descriptor map is empty.")
	}
	sd, ok := allStructDescriptor[sdID]
	if !ok {
		return nil, perros.Errorf("struct ID %s struct descriptor not found.", sdID)
	}
	return sd.Factory, nil
}

// 传入 field 解析出 sdID
func (this *AutowireBase) ParseSDID(field *autowire.FieldInfo) (string, error) {
	return this.sdIDParser.Parse(field)
}

// 解析参数
func (this *AutowireBase) ParseParam(sdID string, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	allStructDescriptor := this.facadeAutowire.GetAllStructDescriptors() // SDMap
	if allStructDescriptor == nil {
		return nil, perros.New("struct descriptor map is empty.") // SDMap is empty
	}
	sd, ok := allStructDescriptor[sdID]
	if !ok {
		return nil, perros.Errorf("struct ID %s struct descriptor not found", sdID) // sdID 不存在 SDMap中
	}
	if sd.ParamFactory == nil {
		return nil, nil
	}
	if sd.ParamLoader != nil {
		// 尝试使用 sd 参数解析
		param, err := sd.ParamLoader.Load(sd, fieldInfo)
		if err == nil {
			return param, nil
		} else {
			// 日志警告，给定的 paramLoad 加载失败，回退到默认值（通过给的 sdID 找到 sd 中的 paramLoad 是错误的）
			logger.Red("[Autowire Base] Load SD %s param with defined sd.ParamLoader error:%s\n"+
				"Try load by autowire %s's default paramLoader", sd.ID(), err, this.facadeAutowire.TagKey())
		}
	}
	// 使用 autowire 定义的加载器（参数解析器）作为备用
	return this.paramLoader.Load(sd, fieldInfo)
}

func (this *AutowireBase) Construct(sdID string, impledPtr, param interface{}) (interface{}, error) {
	allStructDescriptor := this.facadeAutowire.GetAllStructDescriptors()
	if allStructDescriptor == nil {
		return nil, perros.New("struct descriptor map is empty.")
	}
	sd, ok := allStructDescriptor[sdID]
	if !ok {
		return nil, perros.Errorf("struct ID %s struct descriptor not found.", sdID)
	}
	if sd.ConstructFunc != nil { // 有参构造（用户写的）
		return sd.ConstructFunc(impledPtr, param)
	}
	return impledPtr, nil // 默认无参构造
}

// 注入工厂调用后的切点
func (this *AutowireBase) InjectPosition() autowire.InjectPosition {
	return autowire.AfterFactoryCalled
}
