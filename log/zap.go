package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kostya2011/dht-11-home-sensor/config"
	"github.com/kostya2011/dht-11-home-sensor/interfaces"
)

type ZapLogger struct {
	logger *zap.Logger
}

// var (
// 	logger *zap.Logger
// 	once   sync.Once
// )

func NewZapLogger() *ZapLogger {
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

	logger, err := zapConfig.Build()

	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	return &ZapLogger{
		logger: logger,
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

func (zl *ZapLogger) Debug(msg string, fields ...interfaces.LogFields) {
	zl.logger.Debug(msg)
}

func (zl *ZapLogger) Info(msg string, fields ...interfaces.LogFields) {
	zl.logger.Info(msg, zapFields(fields)...)
}

func (zl *ZapLogger) Warn(msg string, fields ...interfaces.LogFields) {
	zl.logger.Warn(msg, zapFields(fields)...)
}

func (zl *ZapLogger) Error(msg string, fields ...interfaces.LogFields) {
	zl.logger.Error(msg, zapFields(fields)...)
}

func (zl *ZapLogger) Fatal(msg string, fields ...interfaces.LogFields) {
	zl.logger.Fatal(msg, zapFields(fields)...)
}

func zapFields(fields []interfaces.LogFields) []zap.Field {
	logFields := make([]zap.Field, 0, len(fields))
	for _, v := range fields {
		logFields = append(logFields, zap.Any(v.Key, v.Value))
	}
	return logFields
}
