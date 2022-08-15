package mocks

import (
	"github.com/stretchr/testify/mock"
	"ioc-go/autowire"
	"testing"
)

// 是 SDIDParser 类型在自动生成的模拟类型
type SDIDParser struct {
	mock.Mock
}

// 创建一个新的 SDIDParser 实例。它还在模拟上注册了 testing.TB 接口和一个理清函数来断言模拟期望
func NewSDIDParser(t testing.TB) *SDIDParser {
	mock := &SDIDParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}

// 提供了一个具有给定字段的模拟函数：fieldInfo
func (_m *SDIDParser) Parse(fieldInfo *autowire.FieldInfo) (string, error) {
	result := _m.Called(fieldInfo)

	var r0 string
	if rf, ok := result.Get(0).(func(*autowire.FieldInfo) string); ok {
		r0 = rf(fieldInfo)
	} else {
		r0 = result.Get(0).(string)
	}

	var r1 error
	if rf, ok := result.Get(1).(func(*autowire.FieldInfo) error); ok {
		r1 = rf(fieldInfo)
	} else {
		r1 = result.Error(1)
	}

	return r0, r1
}
