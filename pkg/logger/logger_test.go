package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewSugardLogger()
	logger.Info("测试")
}