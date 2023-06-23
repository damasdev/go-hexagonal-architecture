package logger

import (
	"github.com/rs/zerolog"
)

type logLevel uint8

var (
	DebugLevel logLevel = 0 // debug
	InfoLevel  logLevel = 1 // info
	WarnLevel  logLevel = 2 // warn
	ErrorLevel logLevel = 3 // error
	PanicLevel logLevel = 4 // panic
)

var levelMapping = map[logLevel]zerolog.Level{
	DebugLevel: zerolog.DebugLevel,
	InfoLevel:  zerolog.InfoLevel,
	WarnLevel:  zerolog.WarnLevel,
	ErrorLevel: zerolog.ErrorLevel,
	PanicLevel: zerolog.PanicLevel,
}

func toLevel(level logLevel) zerolog.Level {
	if level, ok := levelMapping[level]; ok {
		return level
	}
	return zerolog.InfoLevel
}
