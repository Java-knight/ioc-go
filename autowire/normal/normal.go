package normal

import (
	"ioc-go/autowire"
	"ioc-go/autowire/base"
	"ioc-go/autowire/param_loader"
	"ioc-go/autowire/sdid_parse"
)

const Name = "normal"

// Go语言执行顺序：全局声明（变量、常量...）——> init() ——> main()
func init() {
	// 注册 autowire, 将其 wrapper到Map中
	autowire.RegisterAutowire(NewNormalAutowire(nil, nil, nil))
}

// 创建一个基于自动装配到自动装配，例如 config, base,facade 可以重写到外部自动装配
//sp: sdID解析器；pl: 参数加载器；facade: 上层到自动装配
func NewNormalAutowire(sp autowire.SDIDParser, pl autowire.ParamLoader, facade autowire.Autowire) autowire.Autowire {
	if sp == nil {
		sp = sdid_parse.GetDefaultSDIDParser()
	}
	if pl == nil {
		pl = param_loader.GetDefaultParamLoader()
	}
	normalAutowire := &NormalAutowire{}
	if facade == nil {
		facade = normalAutowire
	}
	normalAutowire.AutowireBase = base.New(facade, sp, pl)
	return normalAutowire
}

// 普通的 autowire【非单例】
type NormalAutowire struct {
	base.AutowireBase
}

func (n *NormalAutowire) TagKey() string {
	return Name
}

func (n *NormalAutowire) IsSingleton() bool {
	return false
}

func (n *NormalAutowire) CanBeEntrance() bool {
	return false
}

func (n *NormalAutowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	return normalEntryDescriptorMap
}

var _ (autowire.Autowire) = (*NormalAutowire)(nil)
var normalEntryDescriptorMap = make(map[string]*autowire.StructDescriptor)

// dev APIs

// 注册sd
func RegisterStructDescriptor(sd *autowire.StructDescriptor) {
	sdID := sd.ID()
	normalEntryDescriptorMap[sdID] = sd
	autowire.RegisterStructDescriptor(sdID, sd)
	if sd.Alias != "" {
		autowire.RegisterAlias(sd.Alias, sdID)
	}
}

// 获取实现
func GetImpl(sdID string, param interface{}) (interface{}, error) {
	return autowire.Impl(Name, sdID, param)
}

// 使用代理获取实现
func GetImplWithProxy(sdID string, param interface{}) (interface{}, error) {
	return autowire.ImplWithProxy(Name, sdID, param)
}
