package logger

import (
	"github.com/rs/zerolog"
)

type options struct {
	data *map[string]interface{}
	err  *error
}

type option func(opts *options)

func WithData(data map[string]interface{}) option {
	return func(opts *options) {
		opts.data = &data
	}
}

func WithError(err error) option {
	return func(options *options) {
		options.err = &err
	}
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

	event.Str("service", log.name)

	return event
}

func (opts *options) getData() *map[string]interface{} {
	return opts.data
}

func (opts *options) getError() *error {
	return opts.err
}
