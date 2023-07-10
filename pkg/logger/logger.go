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

func (zeroLog *logger) Debug(message string, opts ...log.OptionFunc) {
	zeroLog.withContext(zeroLog.handler.Debug(), opts...).Msg(message)
}

func (zeroLog *logger) Info(message string, opts ...log.OptionFunc) {
	zeroLog.withContext(zeroLog.handler.Info(), opts...).Msg(message)
}

func (zeroLog *logger) Warning(message string, opts ...log.OptionFunc) {
	zeroLog.withContext(zeroLog.handler.Warn(), opts...).Msg(message)
}

func (zeroLog *logger) Error(message string, opts ...log.OptionFunc) {
	zeroLog.withContext(zeroLog.handler.Error(), opts...).Msg(message)
}

func (zeroLog *logger) Panic(message string, opts ...log.OptionFunc) {
	zeroLog.withContext(zeroLog.handler.Panic(), opts...).Msg(message)
}

func (zeroLog *logger) withContext(event *zerolog.Event, opts ...log.OptionFunc) *zerolog.Event {

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

	event.Str("service", zeroLog.name)

	return event
}
