package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerDisable(t *testing.T) {
	assert.True(t, !disableLogs)
	Disable()
	assert.True(t, disableLogs)
}
