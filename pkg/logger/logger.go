package logger

import (
	"os"

	"github.com/damasdev/fiber/pkg/log"
	"github.com/rs/zerolog"
)

var (
	Logger iLogger
)

type iLogger interface {
	Debug(message string, opts ...log.Option)
	Info(message string, opts ...log.Option)
	Warning(message string, opts ...log.Option)
	Error(message string, opts ...log.Option)
	Panic(message string, opts ...log.Option)
}

type zeroLog struct {
	name    string
	handler zerolog.Logger
}

func Initialize(cfgs ...config) {

	config := &configs{}
	for _, cfg := range cfgs {
		cfg(config)
	}

	if config.writer == nil {
		config.writer = os.Stdout
	}

	handler := zerolog.New(config.writer).With().Timestamp().Logger()

	Logger = &zeroLog{
		name:    *config.name,
		handler: handler.Level(toLevel(*config.level)),
	}
}

func (l *zeroLog) Debug(message string, opts ...log.Option) {
	l.withContext(l.handler.Debug(), opts...).Msg(message)
}

func (l *zeroLog) Info(message string, opts ...log.Option) {
	l.withContext(l.handler.Info(), opts...).Msg(message)
}

func (l *zeroLog) Warning(message string, opts ...log.Option) {
	l.withContext(l.handler.Warn(), opts...).Msg(message)
}

func (l *zeroLog) Error(message string, opts ...log.Option) {
	l.withContext(l.handler.Error(), opts...).Msg(message)
}

func (l *zeroLog) Panic(message string, opts ...log.Option) {
	l.withContext(l.handler.Panic(), opts...).Msg(message)
}

func (l *zeroLog) withContext(event *zerolog.Event, opts ...log.Option) *zerolog.Event {

	log := &log.Options{}
	for _, opt := range opts {
		opt(log)
	}

	if data := log.GetData(); data != nil {
		event.Interface("data", data)
	}

	if err := log.GetError(); err != nil {
		event.Err(*err)
	}

	event.Str("service", l.name)

	return event
}
