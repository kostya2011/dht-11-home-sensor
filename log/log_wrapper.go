package log

import (
	"github.com/kostya2011/home-data/interfaces"
)

var (
	lw interfaces.Logger
)

func Init(logger interfaces.Logger) {
	lw = logger
}

func Debug(msg string, fields ...interfaces.LogFields) {
	lw.Debug(msg, fields...)
}

func Info(msg string, fields ...interfaces.LogFields) {
	lw.Info(msg, fields...)
}

func Warn(msg string, fields ...interfaces.LogFields) {
	lw.Warn(msg, fields...)
}

func Error(msg string, fields ...interfaces.LogFields) {
	lw.Error(msg, fields...)
}

func Fatal(msg string, fields ...interfaces.LogFields) {
	lw.Fatal(msg, fields...)
}

func SetField(key string, value interface{}) interfaces.LogFields {
	return interfaces.LogFields{
		Key:   key,
		Value: value,
	}
}
