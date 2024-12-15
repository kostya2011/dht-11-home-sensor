package log

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kostya2011/dht-11-home-sensor/config"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func init() {
	// var err error
	var err error
	level := zap.NewAtomicLevelAt(getLoggerLevel(config.Cfg.Log.Level))
	encoder := zap.NewProductionEncoderConfig()

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = encoder
	zapConfig.Level = level
	zapConfig.Encoding = config.Cfg.Log.EncondingFormat
	zapConfig.InitialFields = map[string]interface{}{"serverMode": config.Cfg.Server.Mode}
	zapConfig.OutputPaths = config.Cfg.Log.Outputs
	zapConfig.ErrorOutputPaths = config.Cfg.Log.ErrorOutputs

	if config.Cfg.Server.Mode == "production" {
		zapConfig.Development = false
	} else {
		zapConfig.Development = true
	}

	logger, err = zapConfig.Build()

	if err != nil {
		panic(err)
	}
}

func getLoggerLevel(lvl string) zapcore.Level {
	switch lvl {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	}

	return zap.ErrorLevel
}

func IntLogField(key string, value int) zapcore.Field {
	return zap.Int(key, value)
}

func StringLogField(key string, value string) zapcore.Field {
	return zap.String(key, value)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Erorr(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

func GetRequestsLoggerFormat() {
	fmt.Println("Hello")
}
