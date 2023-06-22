package logger

import (
	"github.com/rs/zerolog"
)

type LogLevel uint8

var (
	DebugLevel LogLevel = 0 // debug
	InfoLevel  LogLevel = 1 // info
	WarnLevel  LogLevel = 2 // warn
	ErrorLevel LogLevel = 3 // error
	PanicLevel LogLevel = 4 // panic
)

var levelMapping = map[LogLevel]zerolog.Level{
	DebugLevel: zerolog.DebugLevel,
	InfoLevel:  zerolog.InfoLevel,
	WarnLevel:  zerolog.WarnLevel,
	ErrorLevel: zerolog.ErrorLevel,
	PanicLevel: zerolog.PanicLevel,
}

func toLevel(level LogLevel) zerolog.Level {
	if level, ok := levelMapping[level]; ok {
		return level
	}
	return zerolog.InfoLevel
}
