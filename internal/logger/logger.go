package logger

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	// level
	looger, _ := zap.NewDevelopment()

	return looger
}
