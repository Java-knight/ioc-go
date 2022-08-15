package param_loader

import "ioc-go/autowire"

// 默认参数加载器
type defaultParamLoader struct {
	defaultConfigParamLoader     autowire.ParamLoader // 配置参数加载器
	defaultTagParamLoader        autowire.ParamLoader // 标签参数加载器
	defaultTagPointToParamLoader autowire.ParamLoader // 标签指针加载器
}

// 默认参数加载器单例
var defaultParamLoaderSingleton autowire.ParamLoader

// 获取默认参数加载器
func GetDefaultParamLoader() autowire.ParamLoader {
	if defaultParamLoaderSingleton == nil {
		defaultParamLoaderSingleton = &defaultParamLoader{
			defaultConfigParamLoader:     GetDefaultConfigParamLoader(),
			defaultTagParamLoader:        GetDefaultTagParamLoader(),
			defaultTagPointToParamLoader: GetDefaultTagPointToConfigParamLoader(),
		}
	}
	return defaultParamLoaderSingleton
}

/*
Load 尝试从 3 中类型中加载配置（从严到松）：
（1）尝试使用 defaultTagPointToParamLoader 从标签指向的配置中加载
（2）尝试使用 defaultTagParamLoader 从字段标签加载
（3）尝试使用 defaultConfigParamLoader 从配置参数加载
如果两种方式都失败，它将返回错误
*/
func (this *defaultParamLoader) Load(sd *autowire.StructDescriptor, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	if param, err := this.defaultTagPointToParamLoader.Load(sd, fieldInfo); err == nil {
		return param, err
	}
	// TODO: log warning
	if param, err := this.defaultTagParamLoader.Load(sd, fieldInfo); err == nil {
		return param, err
	}
	// TODO log warning
	return this.defaultConfigParamLoader.Load(sd, fieldInfo)
}
