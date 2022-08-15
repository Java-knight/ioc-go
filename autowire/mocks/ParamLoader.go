package mocks

import (
	"github.com/stretchr/testify/mock"
	"ioc-go/autowire"
	"testing"
)

// ParamLoader 是 参数加载器类型的自动生成的模拟类型
type ParamLoader struct {
	mock.Mock
}

// 创建一个新的 ParamLoader 实例。它还在模拟上注册了 testing.TB 接口和一个清理函数来断言模拟期望值
func NewParamLoader(t testing.TB) *ParamLoader {
	mock := &ParamLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}

// 提供了一个具有给定字段的模拟函数：sd, fieldInfo
func (_m *ParamLoader) Load(sd *autowire.StructDescriptor, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	result := _m.Called(sd, fieldInfo)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(*autowire.StructDescriptor, *autowire.FieldInfo) interface{}); ok {
		r0 = rf(sd, fieldInfo)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(*autowire.StructDescriptor, *autowire.FieldInfo) error); ok {
		r1 = rf(sd, fieldInfo)
	} else {
		r1 = result.Error(1)
	}
	return r0, r1
}
