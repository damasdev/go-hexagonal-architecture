package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type iLogger interface {
	Debug(message string, opts ...option)
	Info(message string, opts ...option)
	Warning(message string, opts ...option)
	Error(message string, opts ...option)
	Fatal(message string, opts ...option)
	Panic(message string, opts ...option)
}

var (
	Logger iLogger
)

type zeroLog struct {
	handler zerolog.Logger
}

func Initialize(writer io.Writer, level LogLevel) {
	handler := zerolog.New(writer).With().Timestamp().Logger()

	Logger = &zeroLog{
		handler: handler.Level(toLevel(level)),
	}
}

func (log *zeroLog) Debug(message string, opts ...option) {
	withContext(log.handler.Debug(), opts...).Msg(message)
}

func (log *zeroLog) Info(message string, opts ...option) {
	withContext(log.handler.Info(), opts...).Msg(message)
}

func (log *zeroLog) Warning(message string, opts ...option) {
	withContext(log.handler.Warn(), opts...).Msg(message)
}

func (log *zeroLog) Error(message string, opts ...option) {
	withContext(log.handler.Error(), opts...).Msg(message)
}

func (log *zeroLog) Fatal(message string, opts ...option) {
	withContext(log.handler.Fatal(), opts...).Msg(message)
}

func (log *zeroLog) Panic(message string, opts ...option) {
	withContext(log.handler.Panic(), opts...).Msg(message)
}

func withContext(event *zerolog.Event, opts ...option) *zerolog.Event {

	model := &options{}
	for _, opt := range opts {
		opt(model)
	}

	if data := model.getData(); data != nil {
		event.Interface("data", data)
	}

	if err := model.getError(); err != nil {
		event.Err(*err)
	}

	return event
}
