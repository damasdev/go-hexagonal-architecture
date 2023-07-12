package logger

import (
	"os"

	"go-hexagonal-architecture/pkg/log"

	"github.com/rs/zerolog"
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
	handler zerolog.Logger
}

func Initialize(cfgs ...configFunc) {

	cfg := &Config{}
	for _, fn := range cfgs {
		fn(cfg)
	}

	if cfg.GetWriter() == nil {
		cfg.writer = os.Stderr
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	handler := zerolog.New(cfg.writer).With().Str("service", cfg.GetName()).Timestamp()

	Logger = &logger{
		handler: handler.Logger().Level(ToLevel(*cfg.GetLevel())),
	}
}

func (l *logger) Debug(message string, opts ...log.OptionFunc) {
	withContext(l.handler.Debug(), opts...).Msg(message)
}

func (l *logger) Info(message string, opts ...log.OptionFunc) {
	withContext(l.handler.Info(), opts...).Msg(message)
}

func (l *logger) Warning(message string, opts ...log.OptionFunc) {
	withContext(l.handler.Warn(), opts...).Msg(message)
}

func (l *logger) Error(message string, opts ...log.OptionFunc) {
	withContext(l.handler.Error(), opts...).Msg(message)
}

func (l *logger) Panic(message string, opts ...log.OptionFunc) {
	withContext(l.handler.Panic(), opts...).Msg(message)
}

func withContext(event *zerolog.Event, opts ...log.OptionFunc) *zerolog.Event {

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

	if skip := opt.GetSkip(); skip > 0 {
		event.Caller(opt.GetSkip())
	} else {
		event.Caller(2)
	}

	return event
}
