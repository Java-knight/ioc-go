package config

import (
	perros "github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// 加载属性(从yaml中读取赋值)【递归调用】
func loadProperty(splitedConfigName []string, index int, tempConfigMap Config, configStructPtr interface{}) error {
	subConfig, ok := tempConfigMap[splitedConfigName[index]]
	if !ok {
		perros.Errorf("property %s's key %s not found", splitedConfigName, splitedConfigName[index])
	}
	if index+1 == len(splitedConfigName) { // yaml的赋值操作
		targetConfigByte, err := yaml.Marshal(subConfig)
		if err != nil {
			return perros.Errorf("property %s's key %s invalid, error = %s", splitedConfigName, splitedConfigName[index], err)
		}
		err = yaml.Unmarshal(targetConfigByte, configStructPtr)
		if err != nil {
			return perros.Errorf("property %s's key %s doesn't match type %+v, error = %s", splitedConfigName, splitedConfigName[index], configStructPtr, err)
		}
		return nil
	}

	subMap, ok := subConfig.(Config)
	if !ok {
		return perros.Errorf("property %s's key %s of config is not map[string]string, which is %+v", splitedConfigName, splitedConfigName[index], subMap)
	}
	return loadProperty(splitedConfigName, index+1, subMap, configStructPtr)
}
