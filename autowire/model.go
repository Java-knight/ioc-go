package autowire

import (
	"ioc-go/autowire/util"
	"reflect"
)

/*
结构描述符(Struct Descriptor, SD):
    用于描述一个被开发者定义的结构，包含对象生命周期的全部信息，例如结构类型是什么，实现了哪些接口，如何被构造等等。
*/
// Autowire
type Autowire interface {
	// 充当包装 WrapperAutowireMap中的key，Map（autowire.TagKey-wrapperAutowire）
	TagKey() string

	// IsSingleton 表示 struct 可以作为启动入口，并且全局只有一个 impl，只创建一次
	IsSingleton() bool

	// TODO: 表示结构在应用程序开始时加载。默认情况下，只有 rpc-server 可以作为入口
	CanBeEntrance() bool

	Factory(sdID string) (interface{}, error)

	// 解析结构性描述符ID（import:struct）
	ParseSDID(field *FieldInfo) (string, error)
	// 将结构性描述符信息 解析成的参数（AST->token）
	ParseParam(sdID string, fieldInfo *FieldInfo) (interface{}, error)
	// 构造函数
	Construct(sdID string, impledPtr, param interface{}) (interface{}, error)
	// 获取全部结构性描述符
	GetAllStructDescriptors() map[string]*StructDescriptor
	InjectPosition() InjectPosition
}

var wrapperAutowireMap = make(map[string]WrapperAutowire)

// 注册 Autowire（将 Wrapper 后的 Autowire 存入 Map 中）
func RegisterAutowire(autowire Autowire) {
	wrapperAutowireMap[autowire.TagKey()] = getWrappedAutowire(autowire, wrapperAutowireMap)
}

// 获取 WrapperAutowireMap
func GetAllWrapperAutowires() map[string]WrapperAutowire {
	return wrapperAutowireMap
}

/*
FieldInfo 介绍
（1）如果是 struct
      MyStruct *MyStruct `autowire-type:"MyStruct"`
      FieldInfo --->
      FieldInfo.FieldName == "MyStruct"
      FieldInfo.FieldType == ""  // 用来区分是 struct 还是 interface
      FieldInfo.TagKey == "autowire-type"
      FieldInfo.TagValue == "MyStruct"

（2）如果是 interface
      MyStruct MyInterface `autowire-type:"MyStruct"`
      FieldInfo --->
      FieldInfo.FieldName == "MyStruct"
      FieldInfo.FieldType == "MyInterface"
      FieldInfo.TagKey == "autowire-type"
      FieldInfo.TagValue == "MyStruct"
*/
// FieldInfo 结构性描述符 ——> 属性信息
type FieldInfo struct {
	FieldName         string
	FieldType         string
	TagKey            string
	TagValue          string
	FieldReflectType  reflect.Type  // 属性类型
	FieldReflectValue reflect.Value // 属性类型的值
}

// StructDescriptor 结构性描述符
type StructDescriptor struct {
	Factory       func() interface{} // 结构的工厂函数，返回值是未经初始化的空结构指针
	ParamFactory  func() interface{} // 参数工厂
	ParamLoader   ParamLoader
	ConstructFunc func(impl interface{}, param interface{}) (interface{}, error) // 注入（构造函数）
	DestroyFunc   func(impl interface{})                                         // 销毁对象过程（非必要）
	Alias         string                                                         // 给结构性描述符的别名，SDID
	DisableProxy  bool                                                           // 是否 禁用代理 和 AOP
	Metadata      Metadata                                                       // SD的元数据Map

	impledStructPtr interface{} // 仅用于获取名称
}

// SDID(包名.结构体/接口名)
func (sd *StructDescriptor) ID() string {
	return util.GetSDIDByStructPtr(sd.getStructPtr())
}

func (sd *StructDescriptor) getStructPtr() interface{} {
	if sd.impledStructPtr == nil { // 表示需要工厂创建
		sd.impledStructPtr = sd.Factory()
	}
	return sd.impledStructPtr
}

// 加载参数的接口
type ParamLoader interface {
	Load(sd *StructDescriptor, fieldInfo *FieldInfo) (interface{}, error)
}

// 解析 SDID 的接口
type SDIDParser interface {
	Parse(field *FieldInfo) (string, error)
}

// Metadata 是 SD 的元数据
type Metadata map[string]interface{}

// 注入点
type InjectPosition int

const (
	AfterFactoryCalled     InjectPosition = 0 // 工厂调用后，切点
	AfterConstructorCalled InjectPosition = 1 // 构造函数后，切点
)
