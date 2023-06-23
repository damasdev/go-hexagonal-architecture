package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type iLogger interface {
	Debug(message string, opts ...option)
	Info(message string, opts ...option)
	Warning(message string, opts ...option)
	Error(message string, opts ...option)
	Panic(message string, opts ...option)

	SetRequestTime()
}

var (
	Logger iLogger
)

type zeroLog struct {
	name        string
	requestTime *time.Time

	handler zerolog.Logger
}

func Initialize(cfgs ...config) {
	handler := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := &configs{}
	for _, opt := range cfgs {
		opt(config)
	}

	Logger = &zeroLog{
		name:        *config.name,
		requestTime: nil,

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

func (log *zeroLog) SetRequestTime() {
	time := time.Now()

	log.requestTime = &time
}
