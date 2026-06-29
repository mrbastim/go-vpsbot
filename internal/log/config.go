package log

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func InitLogger() {
	var err error
	config := zap.NewProductionConfig()
	enccoderConfig := zap.NewProductionEncoderConfig()
	config.EncoderConfig = enccoderConfig

	zapLogger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLogger.Fatal(msg, fields...)
}

func Sync() error {
	return zapLogger.Sync()
}

func GetLogger() *zap.Logger {
	if zapLogger == nil {
		panic("logger not initialized")
	}
	return zapLogger
}
