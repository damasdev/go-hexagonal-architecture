package logger

import (
	"io"
)

type Configs struct {
	name  *string
	level *LogLevel

	writer io.Writer
}

type config func(opts *Configs)

func WithLevel(level LogLevel) config {
	return func(opts *Configs) {
		opts.level = &level
	}
}

func WithName(name string) config {
	return func(configs *Configs) {
		configs.name = &name
	}
}

func WithWriter(writer io.Writer) config {
	return func(configs *Configs) {
		configs.writer = writer
	}
}

func (options *Configs) GetName() *string {
	return options.name
}

func (options *Configs) GetLevel() *LogLevel {
	return options.level
}

func (options *Configs) GetWriter() io.Writer {
	return options.writer
}
