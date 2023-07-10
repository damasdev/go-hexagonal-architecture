package logger

import (
	"os"
	"time"

	"github.com/damasdev/fiber/pkg/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	Logger iLogger
)

type iLogger interface {
	Debug(message string, opts ...log.OptionFunc)
	Info(message string, opts ...log.OptionFunc)
	Warning(message string, opts ...log.OptionFunc)
	Error(message string, opts ...log.OptionFunc)
	Panic(message string, opts ...log.OptionFunc)
}

type logger struct {
	name    string
	handler zerolog.Logger
}

func Initialize(cfgs ...configFunc) {

	cfg := &Config{}
	for _, fn := range cfgs {
		fn(cfg)
	}

	if cfg.writer == nil {
		cfg.writer = os.Stdout
	}

	zerolog.DurationFieldUnit = time.Nanosecond
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	Logger = &logger{
		name:    *cfg.name,
		handler: zerolog.New(cfg.writer).With().Timestamp().Logger().Level(ToLevel(*cfg.level)),
	}
}

func (l *logger) Debug(message string, opts ...log.OptionFunc) {
	l.withContext(l.handler.Debug(), opts...).Msg(message)
}

func (l *logger) Info(message string, opts ...log.OptionFunc) {
	l.withContext(l.handler.Info(), opts...).Msg(message)
}

func (l *logger) Warning(message string, opts ...log.OptionFunc) {
	l.withContext(l.handler.Warn(), opts...).Msg(message)
}

func (l *logger) Error(message string, opts ...log.OptionFunc) {
	l.withContext(l.handler.Error(), opts...).Msg(message)
}

func (l *logger) Panic(message string, opts ...log.OptionFunc) {
	l.withContext(l.handler.Panic(), opts...).Msg(message)
}

func (l *logger) withContext(event *zerolog.Event, opts ...log.OptionFunc) *zerolog.Event {

	opt := &log.Option{}
	for _, fn := range opts {
		fn(opt)
	}

	if data := opt.GetData(); data != nil {
		event.Interface("data", data)
	}

	if err := opt.GetError(); err != nil {
		event.Err(*err)
	}

	event.Str("service", l.name)

	return event
}
