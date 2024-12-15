package interfaces

type LogFields map[string]interface{}

type Logger interface {
	Debug(msg string, fields ...LogFields)
	Info(msg string, fields ...LogFields)
	Warn(msg string, fields ...LogFields)
	Error(msg string, fields ...LogFields)
	Fatal(msg string, fields ...LogFields)
}
