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

func (logger *ZapLogger) Debug(msg string, fields ...interfaces.LogFields) {
	logger.logger.Debug(msg)
}

func (logger *ZapLogger) Info(msg string, fields ...interfaces.LogFields) {
	logger.logger.Info(msg)
}

func (logger *ZapLogger) Warn(msg string, fields ...interfaces.LogFields) {
	logger.logger.Warn(msg)
}

func (logger *ZapLogger) Error(msg string, fields ...interfaces.LogFields) {
	logger.logger.Error(msg)
}

func (logger *ZapLogger) Fatal(msg string, fields ...interfaces.LogFields) {
	logger.logger.Fatal(msg)
}
