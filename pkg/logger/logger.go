package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

type iLogger interface {
	Debug(message string, opts ...option)
	Info(message string, opts ...option)
	Warning(message string, opts ...option)
	Error(message string, opts ...option)
	Fatal(message string, opts ...option)
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

func (log *zeroLog) Fatal(message string, opts ...option) {
	log.withContext(log.handler.Fatal(), opts...).Msg(message)
}

func (log *zeroLog) Panic(message string, opts ...option) {
	log.withContext(log.handler.Panic(), opts...).Msg(message)
}

func (log *zeroLog) SetRequestTime() {
	time := time.Now()

	log.requestTime = &time
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

	if log.requestTime != nil {
		processingTime := time.Since(*log.requestTime).Milliseconds()
		event.Str("latency", strconv.Itoa(int(processingTime))+"ms")

		log.requestTime = nil
	}

	return event
}
