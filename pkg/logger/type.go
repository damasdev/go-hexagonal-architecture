package logger

type iLogger interface {
	Debug(message string, opts ...option)
	Info(message string, opts ...option)
	Warning(message string, opts ...option)
	Error(message string, opts ...option)
	Panic(message string, opts ...option)
}
