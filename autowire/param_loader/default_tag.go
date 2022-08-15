package param_loader

import (
	"encoding/json"
	"github.com/pkg/errors"
	"ioc-go/autowire"
	"ioc-go/config"
	"log"
	"strings"
)

// 默认标签
type defaultTag struct {
}

var defaultTagParamLoaderSingleton autowire.ParamLoader

// 获取默认标签参数加载器
func GetDefaultTagParamLoader() autowire.ParamLoader {
	if defaultTagParamLoaderSingleton == nil {
		defaultTagParamLoaderSingleton = &defaultTag{}
	}
	return defaultTagParamLoaderSingleton
}

/*
Load 支持加载参数:
```go
type Config struct {
    Address  string,
    Password string,
    DB       string,
}
```

from field:

```go
NormalRedis normalRedis.Redis `normal:"Impl,address=127.0.0.1&password=xxx&db=0"`
```
*/
func (this *defaultTag) Load(sd *autowire.StructDescriptor, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	if sd == nil || fieldInfo == nil || sd.ParamFactory == nil {
		return nil, errors.New("not supported")
	}
	splitedTagValue := strings.Split(fieldInfo.TagValue, ",")
	if len(splitedTagValue) < 2 {
		return nil, errors.New("not supported")
	}
	kvs := strings.Split(splitedTagValue[1], "&") // {key1-value1, key2-value2, ...}[address-127.0.0.1]
	kvMap := make(map[string]interface{})
	for _, kv := range kvs {
		splitedKV := strings.Split(kv, "=")
		if len(splitedKV) != 2 {
			return nil, errors.New("not supported")
		}
		expandValue := config.ExpandConfigValueIfNecessary(splitedKV[1])
		kvMap[splitedKV[0]] = expandValue
	}

	// 将kvMap 转化为 ParamFactory类型（func() interface{}）
	data, err := json.Marshal(kvMap)
	if err != nil {
		log.Printf("error json marshal %s\n", err)
		return nil, err
	}
	param := sd.ParamFactory()
	if err := json.Unmarshal(data, param); err != nil {
		log.Printf("error json unmashal %s\n", err)
		return nil, err
	}
	return param, nil
}
