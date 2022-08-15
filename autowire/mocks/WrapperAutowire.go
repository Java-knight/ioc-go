package mocks

import (
	"github.com/stretchr/testify/mock"
	"ioc-go/autowire"
	"testing"
)

var _ autowire.Autowire = (*WrapperAutowire)(nil)

// 是 WrapperAutowire 类型的自动生成的模拟类型
type WrapperAutowire struct {
	mock.Mock
}

// 创建 NewWrapperAutowire 的新实例。它还在模拟上注册了 testing.TB 接口和一个清理函数来断言模拟期望
func NewWrapperAutowire(t testing.TB) *WrapperAutowire {
	mock := &WrapperAutowire{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}

// 提供了一个具有给定字段的函数模拟（入口）
func (this *WrapperAutowire) CanBeEntrance() bool {
	result := this.Called()

	var r0 bool
	if rf, ok := result.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = result.Get(0).(bool)
	}

	return r0
}

// 提供了一个具有给定字段的模拟函数：sdID
func (this *WrapperAutowire) Factory(sdID string) (interface{}, error) {
	result := this.Called(sdID)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(string) interface{}); ok {
		r0 = rf(sdID)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(string) error); ok {
		r1 = rf(sdID)
	} else {
		r1 = result.Error(1).(error)
	}
	return r0, r1
}

// 提供一个具有给定字段的模拟函数
func (this *WrapperAutowire) TagKey() string {
	result := this.Called()

	var r0 string
	if rf, ok := result.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = result.Get(0).(string)
	}

	return r0
}

// 提供了一个具有给定字段的模拟函数
func (this *WrapperAutowire) IsSingleton() bool {
	result := this.Called()

	var r0 bool
	if rf, ok := result.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = result.Get(0).(bool)
	}

	return r0
}

// 提供了一个具有给定字段的模拟函数：field
func (this *WrapperAutowire) ParseSDID(field *autowire.FieldInfo) (string, error) {
	result := this.Called(field)

	var r0 string
	if rf, ok := result.Get(0).(func(*autowire.FieldInfo) string); ok {
		r0 = rf(field)
	} else {
		r0 = result.Get(0).(string)
	}

	var r1 error
	if rf, ok := result.Get(1).(func(*autowire.FieldInfo) error); ok {
		r1 = rf(field)
	} else {
		r1 = result.Error(1)

	}

	return r0, r1
}

// 提供了一个具有给定字段的模拟函数：sdID、fieldInfo
func (this *WrapperAutowire) ParseParam(sdID string, fieldInfo *autowire.FieldInfo) (interface{}, error) {
	result := this.Called(sdID, fieldInfo)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(string, *autowire.FieldInfo) interface{}); ok {
		r0 = rf(sdID, fieldInfo)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(string, *autowire.FieldInfo) error); ok {
		r1 = rf(sdID, fieldInfo)
	} else {
		r1 = result.Error(1)
	}

	return r0, r1
}

// 提供具有给定字段的模拟函数：sdID、impledPtr、param
func (this *WrapperAutowire) Construct(sdID string, impledPtr, param interface{}) (interface{}, error) {
	result := this.Called(sdID, impledPtr, param)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(string, interface{}, interface{}) interface{}); ok {
		r0 = rf(sdID, impledPtr, param)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(string, interface{}, interface{}) error); ok {
		r1 = rf(sdID, impledPtr, param)
	} else {
		r1 = result.Error(1)
	}

	return r0, r1
}

// 提供了一个具有给定字段的模拟函数
func (this *WrapperAutowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	result := this.Called()

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

// 提供了一个具有给定字段的模拟函数
func (this *WrapperAutowire) InjectPosition() autowire.InjectPosition {
	result := this.Called()

	var r0 autowire.InjectPosition
	if rf, ok := result.Get(0).(func() autowire.InjectPosition); ok {
		r0 = rf()
	} else {
		r0 = result.Get(0).(autowire.InjectPosition)
	}

	return r0
}

// 提供了一个具有给定字段的模拟函数: sdID、param
func (this *WrapperAutowire) ImplWithParam(sdID string, param interface{}) (interface{}, error) {
	result := this.Called(sdID, param)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(sdID, param)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(sdID, param)
	} else {
		r1 = result.Error(1)
	}
	return r0, r1
}

// 提供了一个具有给定字段的模拟函数: sdID
func (this *WrapperAutowire) ImplWithoutParam(sdID string) (interface{}, error) {
	result := this.Called(sdID)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(string) interface{}); ok {
		r0 = rf(sdID)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(string) error); ok {
		r1 = rf(sdID)
	} else {
		r1 = result.Error(1)
	}
	return r0, r1
}

// 提供了一个具有给定字段的模拟函数: info
func (this *WrapperAutowire) implWithField(info *autowire.FieldInfo) (interface{}, error) {
	result := this.Called(info)

	var r0 interface{}
	if rf, ok := result.Get(0).(func(*autowire.FieldInfo) interface{}); ok {
		r0 = rf(info)
	} else {
		if result.Get(0) != nil {
			r0 = result.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := result.Get(1).(func(*autowire.FieldInfo) error); ok {
		r1 = rf(info)
	} else {
		r1 = result.Error(1)
	}

	return r0, r1
}
