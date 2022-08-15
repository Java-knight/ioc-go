package config

import "regexp"

// 是否是环境变量
func isEnv(envValue string) bool {
	// ${REDIS_ADDRESS_EXPAND}
	ok, err := regexp.Match("^\\$\\{[A-Z_]+}$", []byte(envValue))
	if err != nil || !ok {
		return false
	}
	return true
}
