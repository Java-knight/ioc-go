package singleton

import (
	"ioc-go/autowire"
	"ioc-go/autowire/base"
	"ioc-go/autowire/param_loader"
	"ioc-go/autowire/sdid_parse"
)

func init() {
	autowire.RegisterAutowire(NewSingletonAutowire(nil, nil, nil))
}

const Name = "singleton"

// autowire APIs

// NewSingletonAutowire 创建一个基于单例自动装配(autowire)的 autowire; 例如 grpc、base.facade 可以修改外部其他 autowire
func NewSingletonAutowire(sp autowire.SDIDParser, pl autowire.ParamLoader, facade autowire.Autowire) autowire.Autowire {
	if sp == nil {
		sp = sdid_parse.GetDefaultSDIDParser()
	}
	if pl == nil {
		pl = param_loader.GetDefaultParamLoader()
	}
	singletonAutowire := &SingletonAutowire{
		paramLoader: pl,
		sdIDParser:  sp,
	}
	// 设置入口
	if facade == nil {
		facade = singletonAutowire
	}
	singletonAutowire.AutowireBase = base.New(facade, sp, pl)
	return singletonAutowire
}

type SingletonAutowire struct {
	base.AutowireBase
	paramLoader autowire.ParamLoader
	sdIDParser  autowire.SDIDParser
}

func (s *SingletonAutowire) TagKey() string {
	return Name
}

func (s *SingletonAutowire) IsSingleton() bool {
	return true
}

func (s *SingletonAutowire) CanBeEntrance() bool {
	return false
}

func (s *SingletonAutowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	return singletonStructDescriptorsMap
}

var _ (autowire.Autowire) = (*SingletonAutowire)(nil)
var singletonStructDescriptorsMap = make(map[string]*autowire.StructDescriptor)

// dev APIs

// 注册singleton sd
func RegisterStructDescriptor(sd *autowire.StructDescriptor) {
	sdID := sd.ID()
	singletonStructDescriptorsMap[sdID] = sd
	autowire.RegisterStructDescriptor(sdID, sd)
	if sd.Alias != "" {
		autowire.RegisterAlias(sd.Alias, sdID)
	}
}

// 获取实现
func GetImpl(sdId string, param interface{}) (interface{}, error) {
	return autowire.Impl(Name, sdId, param)
}

// 使用代理获取实现
func GetImplWithProxy(sdId string, param interface{}) (interface{}, error) {
	return autowire.ImplWithProxy(Name, sdId, param)
}
