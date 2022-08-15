package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	EnvPrefixKey = "${"
	EnvSuffixKey = "}"
	And          = "&"
)

// 展开配置环境变量
func ExpandConfigEnvValue(targetValue interface{}) (interface{}, bool) {
	if tv, ok := targetValue.(string); ok {
		// ${REDIS_ADDRESS_EXPAND}, 并且在必须得是env变量
		if strings.HasPrefix(tv, EnvPrefixKey) && strings.HasSuffix(tv, EnvSuffixKey) && isEnv(tv) {
			expandValue := os.ExpandEnv(tv)
			if expandValue != "" { // 在yaml中找到了这个env变量信息，返回真正到值（yaml赋的值）
				return expandValue, true
			}
		}
	}
	return targetValue, false
}

// 展开配置嵌套值
func ExpandConfigNestedValue(targetValue interface{}) (interface{}, bool) {
	if tv, ok := targetValue.(string); ok {
		if strings.HasPrefix(tv, EnvPrefixKey) && strings.HasSuffix(tv, EnvSuffixKey) && !isEnv(tv) {
			// 尝试嵌套解析
			var nestedValue interface{}
			// ${autowire.normal.<dao.Store>.redis-1.address}
			err := LoadConfigByPrefix(tv[2:len(tv)-1], &nestedValue)
			if err != nil {
				return nestedValue, false
			}
			return nestedValue, true
		}
	}
	return targetValue, false
}

// 如果有必要展开配置
func ExpandConfigValueIfNecessary(targetValue interface{}) interface{} {
	result, _ := ExpandConfigEnvValue(targetValue)
	result, _ = ExpandConfigNestedValue(result)
	return result
}

// 如需要展开
func expandIfNecessary(targetValue string) string {
	// address=${REDIS_ADDRESS_EXPAND}&db=5
	if strings.Contains(targetValue, EnvPrefixKey) && strings.Contains(targetValue, EnvSuffixKey) {
		kvs := strings.Split(targetValue, And)
		kvz := make([]string, 0, len(kvs)) // 开辟了一个和 kvs 相同大小的[]string
		for _, kv := range kvs {
			splitedKV := strings.Split(kv, "=")
			if len(splitedKV) != 2 { // db=5, address=${REDIS_ADDRESS_EXPAND}
				kvz = append(kvz, kv)
				continue
			}
			// 正常情况：db=5, address=${REDIS_ADDRESS_EXPAND}
			subKey := splitedKV[0]
			expandValue := ExpandConfigValueIfNecessary(splitedKV[1]) // 看看 ${REDIS_ADDRESS_EXPAND} 是否需要赋值，从yaml读取到赋值信息
			kvz = append(kvz, fmt.Sprintf("%s=%s", subKey, expandValue))
		}
		targetValue = strings.Join(kvz, And)
	}
	return targetValue
}
