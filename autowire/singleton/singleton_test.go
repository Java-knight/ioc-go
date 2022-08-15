package singleton

import (
	"github.com/stretchr/testify/assert"
	"ioc-go/autowire"
	"testing"
)

type mockImpl struct {
}

const mockImplFullName = "ioc-go/autowire/singleton.mockImpl"

func TestAutowire_RegisterAndGetAllStructDescriptors(t *testing.T) {
	t.Run("test singleton autowire register and get all struct descriptors", func(t *testing.T) {
		sd := &autowire.StructDescriptor{
			Factory: func() interface{} {
				return &mockImpl{}
			},
		}

		RegisterStructDescriptor(sd)
		autowire := &SingletonAutowire{}
		allStructDesc := autowire.GetAllStructDescriptors()
		assert.NotNil(t, allStructDesc)
		structDesc, ok := allStructDesc[mockImplFullName]
		assert.True(t, ok)
		assert.Equal(t, mockImplFullName, structDesc.ID())
	})
}

func TestSingletonAutowire_TagKey(t *testing.T) {
	t.Run("test singleton autowire tag", func(t *testing.T) {
		autowire := &SingletonAutowire{}
		assert.Equal(t, Name, autowire.TagKey())
	})
}

func TestSingletonAutowire_IsSingleton(t *testing.T) {
	t.Run("test singleton autowire isSingleton", func(t *testing.T) {
		autowire := &SingletonAutowire{}
		assert.Equal(t, true, autowire.IsSingleton())
	})
}

func TestSingletonAutowire_CanBeEntrance(t *testing.T) {
	t.Run("test singleton autowire can't be entrance", func(t *testing.T) {
		autowire := &SingletonAutowire{}
		assert.Equal(t, false, autowire.CanBeEntrance())
	})
}

func TestNewSingletonAutowire(t *testing.T) {
	singletonAutowire := NewSingletonAutowire(nil, nil, nil)
	assert.NotNil(t, singletonAutowire)
}
