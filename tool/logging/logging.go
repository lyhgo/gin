package logging

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var lock sync.Mutex

func SetUp() {

	var err error
	if logger == nil {
		lock.Lock()
		if logger == nil {
			zapConfig := zap.NewProductionConfig()
			productionEncoderConfig := zap.NewProductionEncoderConfig()
			productionEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)

			zapConfig.EncoderConfig = productionEncoderConfig
			logger, err = zapConfig.Build()
			if err != nil {
				panic(err)
			}
		}

		lock.Unlock()
	}
}

func GetLogger() *zap.Logger {

	var err error
	if logger == nil {
		lock.Lock()
		if logger == nil {
			zapConfig := zap.NewProductionConfig()

			productionEncoderConfig := zap.NewProductionEncoderConfig()
			productionEncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

			zapConfig.EncoderConfig = productionEncoderConfig
			logger, err = zapConfig.Build()
			if err != nil {
				panic(err)
			}
		}

		lock.Unlock()
	}

	return logger
}

// func Info(msg string, fields ...zapcore.Field) {
// 	defer logger.Sync()
// 	logger.Info(msg, fields...)
// }

// func Warn(msg string, fields ...zapcore.Field) {
// 	defer logger.Sync()
// 	logger.Warn(msg, fields...)
// }

// func Debug(msg string, fields ...zapcore.Field) {
// 	defer logger.Sync()
// 	logger.Debug(msg, fields...)
// }

// func Error(msg string, fields ...zapcore.Field) {
// 	defer logger.Sync()
// 	logger.Error(msg, fields...)
// }
