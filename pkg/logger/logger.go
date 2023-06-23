package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	Logger iLogger
)

type zeroLog struct {
	name    string
	handler zerolog.Logger
}

func Initialize(cfgs ...config) {
	handler := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := &configs{}
	for _, opt := range cfgs {
		opt(config)
	}

	Logger = &zeroLog{
		name:    *config.name,
		handler: handler.Level(toLevel(*config.level)),
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

func (log *zeroLog) Panic(message string, opts ...option) {
	log.withContext(log.handler.Panic(), opts...).Msg(message)
}
