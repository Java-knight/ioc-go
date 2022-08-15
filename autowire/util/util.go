package util

import (
	"fmt"
	"reflect"
)

// 获取 struct name
func GetStructName(v interface{}) string {
	if v == nil {
		return ""
	}
	typeOfInterface := GetTypeFromInterface(v)
	return typeOfInterface.Name()
}

// 通过 结构体/接口指针 获取它的 SDID（包名.结构体/接口名）
func GetSDIDByStructPtr(v interface{}) string {
	if v == nil {
		return ""
	}
	typeOfInterface := GetTypeFromInterface(v)
	return fmt.Sprintf("%s.%s", typeOfInterface.PkgPath(), typeOfInterface.Name())
}

// 获取一个接口的类型
//（1）先把接口转成另一个接口（类似拷贝，对象地址不同）；（2）获取它的值或地址（3）返回上一步的类型
func GetTypeFromInterface(v interface{}) reflect.Type {
	valueOfInterface := reflect.ValueOf(v)
	valueOfElemInterface := valueOfInterface.Elem()
	return valueOfElemInterface.Type()
}

// 是否是指针
func IsPointerField(fieldType reflect.Type) bool {
	return fieldType.Kind() == reflect.Ptr
}

// 是否是切片
func IsSliceField(fieldType reflect.Type) bool {
	return fieldType.Kind() == reflect.Slice
}
