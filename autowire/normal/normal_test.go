package normal

import (
	"github.com/stretchr/testify/assert"
	"ioc-go/autowire"
	"testing"
)

type mockImpl struct {
}

const mockImplName = "ioc-go/autowire/normal.mockImpl"

func TestAutowire_RegisterAndGetAllStructDescriptors(t *testing.T) {
	t.Run("test normal autowire register and get all struct descriptor", func(t *testing.T) {
		sd := &autowire.StructDescriptor{
			Factory: func() interface{} {
				return &mockImpl{}
			},
		}
		RegisterStructDescriptor(sd)
		autowire := &NormalAutowire{}
		allStructDesc := autowire.GetAllStructDescriptors()
		assert.NotNil(t, allStructDesc)
		sdid := mockImplName
		structDesc, ok := allStructDesc[sdid]
		assert.True(t, ok)
		assert.Equal(t, sdid, structDesc.ID())
	})
}

func TestAutowire_TagKey(t *testing.T) {
	t.Run("test normal autowire tag", func(t *testing.T) {
		autowire := &NormalAutowire{}
		assert.Equal(t, Name, autowire.TagKey())
	})
}

func TestAutowire_IsSingleton(t *testing.T) {
	t.Run("test normal autowire isSingleton", func(t *testing.T) {
		autowire := &NormalAutowire{}
		assert.Equal(t, false, autowire.IsSingleton())
	})
}

func TestAutowire_CanBeEntrance(t *testing.T) {
	t.Run("test normal autowire can be entrance", func(t *testing.T) {
		autowire := &NormalAutowire{}
		assert.Equal(t, false, autowire.CanBeEntrance())
	})
}

func TestNewNormalAutowire(t *testing.T) {
	normalAutowore := NewNormalAutowire(nil, nil, nil)
	assert.NotNil(t, normalAutowore)
}
