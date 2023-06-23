package logger

import (
	"os"

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
	name    string
}

func Initialize(cfgs ...config) {
	handler := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := &configs{}
	for _, opt := range cfgs {
		opt(config)
	}

	Logger = &zeroLog{
		handler: handler.Level(toLevel(*config.level)),
		name:    *config.name,
	}
}

func (log *zeroLog) Debug(message string, opts ...option) {
	log.withContext(log.handler.Debug(), opts...).Msg(message)
}

func (log *zeroLog) Info(message string, opts ...option) {
	log.withContext(log.handler.Info(), opts...).Msg(message)
}

func (log *zeroLog) Warning(message string, opts ...option) {
	log.withContext(log.handler.Warn(), opts...).Msg(message)
}

func (log *zeroLog) Error(message string, opts ...option) {
	log.withContext(log.handler.Error(), opts...).Msg(message)
}

func (log *zeroLog) Fatal(message string, opts ...option) {
	log.withContext(log.handler.Fatal(), opts...).Msg(message)
}

func (log *zeroLog) Panic(message string, opts ...option) {
	log.withContext(log.handler.Panic(), opts...).Msg(message)
}

func (log *zeroLog) withContext(event *zerolog.Event, opts ...option) *zerolog.Event {

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

	if log.name != "" {
		event.Str("service", log.name)
	}

	return event
}
