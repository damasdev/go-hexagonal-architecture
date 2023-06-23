package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type iLogger interface {
	Debug(message string, options ...any)
	Info(message string, options ...any)
	Warning(message string, options ...any)
	Error(message string, options ...any)
	Fatal(message string, options ...any)
	Panic(message string, options ...any)
}

var (
	Logger iLogger
)

type zeroLog struct {
	handler zerolog.Logger
}

func New(writer io.Writer, level logLevel) iLogger {
	handler := zerolog.New(writer).With().Timestamp().Logger().Level(toLevel(level))

	Logger = &zeroLog{
		handler: handler,
	}

	return Logger
}

func (log *zeroLog) Debug(message string, options ...any) {
	withContext(log.handler.Debug(), options).Caller().Msg(message)
}

func (log *zeroLog) Info(message string, options ...any) {
	withContext(log.handler.Info(), options).Msg(message)
}

func (log *zeroLog) Warning(message string, options ...any) {
	withContext(log.handler.Warn(), options).Msg(message)
}

func (log *zeroLog) Error(message string, options ...any) {
	withContext(log.handler.Error(), options).Msg(message)
}

func (log *zeroLog) Fatal(message string, options ...any) {
	withContext(log.handler.Fatal(), options).Msg(message)
}

func (log *zeroLog) Panic(message string, options ...any) {
	withContext(log.handler.Panic(), options).Msg(message)
}

func withContext(event *zerolog.Event, options []any) *zerolog.Event {
	for _, opt := range options {
		if err, ok := opt.(error); ok {
			event.Err(err)
		} else {
			event.Interface("data", opt)
		}
	}

	return event
}
