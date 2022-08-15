package config

import "strings"

type Config AnyMap

const (
	YmlExtension            = "yml"
	YamlExtension           = "yaml"
	DefaultSearchConfigName = "config"
	DefaultSearchConfigType = YamlExtension // yaml

	YamlConfigSeparator = "."
)

var (
	config              Config
	supportedConfigType = []string{YmlExtension, YamlExtension}
	DefaultSearchPath   = []string{".", "./config", "./configs"}
)

// 通过perfux加载配置：前缀类似于 'a.b.c' 或 'a.b.<github.comxxx.Impl>.c'，configStructPtr 是接口 ptr
func LoadConfigByPrefix(prefix string, configStructPtr interface{}) error {
	if configStructPtr == nil {
		return nil
	}
	realConfigProperties := make([]string, 0)
	for _, val := range splitPrefix2Units(prefix) {
		if val != "" {
			val = expandIfNecessary(val) // 将每个属性都放入 string[]（里面都元素都是原子的）
			realConfigProperties = append(realConfigProperties, val)
		}
	}
	return loadProperty(realConfigProperties, 0, config, configStructPtr) // 读取yaml进行赋值
}

// 将 prefix 拆分成单元化（最小单位），这个方法需要进行test
func splitPrefix2Units(prefix string) []string {
	configProperties := make([]string, 0)
	if prefix == "" {
		return configProperties
	}
	splited := strings.Split(prefix, "<")
	if len(splited) == 1 { // 其中不含有 '<'
		configProperties = strings.Split(prefix, YamlConfigSeparator)
	} else {
		if splited[0] != "" {
			configProperties = strings.Split(splited[0], YamlConfigSeparator)
		}
		backSpliedList := strings.Split(splited[1], ">")
		configProperties = append(configProperties, backSpliedList[0])
		if backSpliedList[1] != "" {
			configProperties = append(configProperties, strings.Split(strings.TrimPrefix(backSpliedList[1], YamlConfigSeparator), YamlConfigSeparator)...)
		}
	}
	return configProperties
}
