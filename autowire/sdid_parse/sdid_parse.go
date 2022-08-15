package sdid_parse

import (
	"ioc-go/autowire"
	"ioc-go/autowire/util"
	"strings"
)

// SDID 解析
type defaultSDIDParse struct {
}

// 默认SDID解析单例
var defaultSDIDParseSingleton autowire.SDIDParser

// 获取默认的 SDID 解析器
func GetDefaultSDIDParser() autowire.SDIDParser {
	if defaultSDIDParseSingleton == nil {
		defaultSDIDParseSingleton = &defaultSDIDParse{}
	}
	return defaultSDIDParseSingleton
}

func (this *defaultSDIDParse) Parse(field *autowire.FieldInfo) (string, error) {
	injectStructName := field.FieldType
	splitedTagValue := strings.Split(field.TagValue, ",")
	if len(splitedTagValue) > 0 && splitedTagValue[0] != "" {
		injectStructName = splitedTagValue[0]
	} else {
		// 标签中没有结构体的 sdid（field的属性信息不是指针，并且interface 类型名字是以 IOCInterface 结尾的，应该是判断是否是自动生成的代码）
		if !util.IsPointerField(field.FieldReflectType) && strings.HasSuffix(injectStructName, "IOCInterface") {
			// 是接口字段，没有来自标记值有效的 sdid, 并且没有带 IOCInterface 后缀
			// 如果是以 IOCInterface 结尾的，将其后缀修剪后返回（比如IStoreIOCInterface 修剪后 IStore）。赋值给 注入结构体名
			injectStructName = strings.TrimSuffix(field.FieldType, "IOCInterface")
		}
	}
	return autowire.GetSDIDByAliasIfNecessary(injectStructName), nil
}
