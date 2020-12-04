package logger

import (
	"go.uber.org/zap/zapcore"
)

func Info(mgs string, fields ...zapcore.Field) {
	AppLogger.Info(mgs, fields...)
}

func Debug(mgs string, fields ...zapcore.Field) {
	AppLogger.Debug(mgs, fields...)
}

func Warn(mgs string, fields ...zapcore.Field) {
	AppLogger.Warn(mgs, fields...)
}

func Error(mgs string, fields ...zapcore.Field) {
	AppLogger.Error(mgs, fields...)
}

func Sync() {
	_ = AppLogger.Sync()
	_ = GLogger.Sync()
	_ = GinLogger.Sync()
}
