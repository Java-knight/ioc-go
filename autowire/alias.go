package autowire

import "fmt"

const (
	AliasKey = "alias"
)

// SDID 别名Map<别名-SDID>
var sdIDAliasMap = make(map[string]string)

// 注册别名
func RegisterAlias(alias, value string) {
	if _, ok := sdIDAliasMap[alias]; ok {
		// 	别名已经被注册了
		panic(fmt.Sprintf("[Autowire Alias] Duplicate alias:[%s]", alias))
	}
	sdIDAliasMap[alias] = value
}

// 必须通过 别名 获取 SDID
// 逻辑：如果key是别名，就返回真正的 sdid；如果key不是别名，就返回key（key就是真正的sdid，没有别名）
func GetSDIDByAliasIfNecessary(key string) string {
	if mappingSDID, ok := GetSDIDByAlias(key); ok {
		return mappingSDID
	}
	return key
}

// 通过 别名 获取 SDID
func GetSDIDByAlias(alias string) (string, bool) {
	sdid, ok := sdIDAliasMap[alias]
	return sdid, ok
}
