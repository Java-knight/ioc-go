package param_loader

import (
	"fmt"
	"github.com/pkg/errors"
	"ioc-go/autowire"
	"ioc-go/autowire/util"
	"ioc-go/config"
)

// 默认配置
type defaultConfig struct {
}

var defaultConfigParamLoaderSingleton autowire.ParamLoader

// 获取默认配置参数加载器
func GetDefaultConfigParamLoader() autowire.ParamLoader {
	if defaultParamLoaderSingleton == nil {
		defaultConfigParamLoaderSingleton = &defaultConfig{}
	}
	return defaultParamLoaderSingleton
}

/*
Load 加载支持加载struct described 如下:

```go
normal.RegisterStructDescriptor(&autowire.StructDescriptor {
    Factory:  func() interface{} {
        return &Impl{}
    }
    ParamFactory: func() interface{} {
        return &Config{}
    }
    ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
        return i, nil
    },
})

type Config struct {
    Address  string,
    Password string,
    DB       string,
}
```

with
Autowire Type 'normal'
StructName 'Impl'

from:
```yaml
autowire:
  normal:
    github.com/ioc-go/test.Impl:
      param:
        address: 127.0.0.1
        password: xxx
        db: 0
```
*/
func (d defaultConfig) Load(sd *autowire.StructDescriptor, fieldInfo *autowire.FieldInfo) (interface{}, error) {

	if sd == nil || sd.ParamFactory == nil {
		return nil, errors.New("not supported")
	}
	param := sd.ParamFactory()
	prefix := getDefaultConfigPrefix(fieldInfo.TagKey, sd)
	if err := config.LoadConfigByPrefix(prefix, param); err != nil {
		return nil, err
	}
	return param, nil
}

// 获取默认配置前缀
func getDefaultConfigPrefix(autowireType string, sd *autowire.StructDescriptor) string {
	structConfigPathKey := sd.Alias
	if structConfigPathKey == "" {
		structConfigPathKey = util.GetSDIDByStructPtr(sd.Factory())
	}
	// autowireType = 'autowire-type'; structConfigPathKey = 'dao.Store'(SDID);
	// autowire.autowire-type.<dao.Store>.param
	return fmt.Sprintf("autowire%[1]s%[2]s%[1]s<%[3]s>%[1]sparam", config.YamlConfigSeparator, autowireType, structConfigPathKey)
}
