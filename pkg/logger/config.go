package logger

import (
	"io"
)

type configs struct {
	name  *string
	level *LogLevel

	writer io.Writer
}

type config func(opts *configs)

func WithLevel(level LogLevel) config {
	return func(opts *configs) {
		opts.level = &level
	}
}

func WithName(name string) config {
	return func(configs *configs) {
		configs.name = &name
	}
}

func WithWriter(writer io.Writer) config {
	return func(configs *configs) {
		configs.writer = writer
	}
}
