package param_loader

import (
	"fmt"
	"github.com/pkg/errors"
	"ioc-go/autowire"
	"ioc-go/autowire/util"
	"ioc-go/config"
	"strings"
)

// 默认标签指向配置
type defaultTagPointToConfig struct {
}

// 默认标签指向配置单例
var defaultTagPointToConfigSingleton autowire.ParamLoader

// 获取 默认标签——>配置参数加载器
func GetDefaultTagPointToConfigParamLoader() autowire.ParamLoader {
	if defaultTagPointToConfigSingleton == nil {
		defaultTagPointToConfigSingleton = &defaultTagPointToConfig{}
	}
	return defaultTagPointToConfigSingleton
}

/*
Load 支持加载 struct described:
```go
normal.RegisterStructDescriptor(&autowire.StructDescriptor {
    Factory: func interface{} {
        return &Impl{}
    },
    ParamFactory: func interface{} {
        return &Config{}
    },
    ConstructFunc: func(i interface{}, p interface{}) (interface{}. error) {
        return i, nil
    },
})

type Config struct {
    Address  string
    Password string
    DB       string
}
```

with
Autowire type 'normal'
StructName 'Impl'
Field:
    MyRedis Redis `normal:"Impl, redis-1"`

from:

```yaml
extension:
  normal:
    github.com/ioc-go/test.Impl:
      redis-1:
        param:
         address: 127.0.0.1
         password: xxx
         db: 0
```
*/
func (this *defaultTagPointToConfig) Load(sd *autowire.StructDescriptor, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	if fieldInfo == nil || sd == nil || sd.ParamFactory == nil {
		return nil, errors.New("not supported")
	}

	param := sd.ParamFactory()
	splitedTagValue := strings.Split(fieldInfo.TagValue, ",")
	if len(splitedTagValue) < 2 {
		return nil, errors.New("tag value not supported")
	}
	prefix := getdefaultTagPointToConfigPrefix(fieldInfo.TagKey, sd, splitedTagValue[1])
	if err := config.LoadConfigByPrefix(prefix, param); err != nil {
		return nil, err
	}
	return param, nil
}

// 获取 默认标签指向配置前缀
func getdefaultTagPointToConfigPrefix(autowireType string, sd *autowire.StructDescriptor, instanceName string) string {
	pointToKey := sd.Alias
	if pointToKey == "" {
		pointToKey = util.GetSDIDByStructPtr(sd.Factory)
	}
	// 变量：MyRedis Redis `normal:"Impl, redis-1"`
	// 附加类型：autowireType = 'normal'; structConfigPathKey = 'dao.Store'(SDID);
	// resultString: autowire.normal.<dao.Store>.redis-1.param
	return fmt.Sprintf("autowire%[1]s%[2]s%[1]s<%[3]s>%[1]s%[4]s%[1]sparam", config.YamlConfigSeparator, autowireType, pointToKey, instanceName)
}
