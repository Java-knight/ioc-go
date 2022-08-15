package mocks

import (
	"github.com/stretchr/testify/mock"
	"ioc-go/autowire"
	"testing"
)

// FacadeAutowire 是 FacadeAutowire 类型的自动生成的模拟（模版）类型
type FacadeAutowire struct {
	mock.Mock // 生成模拟数据
}

// 创建一个新的 NewFacadeAutowire 实例。它还在模拟上注册了 testing.TB 接口和一个清理函数来断言模拟期望
func NewFacadeAutowire(t testing.TB) *FacadeAutowire {
	mock := &FacadeAutowire{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}

// 获取 SDMap
func (_m *FacadeAutowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	result := _m.Called()

	var r0 map[string]*autowire.StructDescriptor
	if rf, ok := result.Get(0).(func() map[string]*autowire.StructDescriptor); ok {
		r0 = rf()
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(map[string]*autowire.StructDescriptor)
		}
	}
	return r0
}

// TagKey 提供了一个具有给定字段的模拟函数
func (_m *FacadeAutowire) TagKey() string {
	result := _m.Called()

	var r0 string
	if rf, ok := result.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = result.Get(0).(string)
	}
	return r0
}
