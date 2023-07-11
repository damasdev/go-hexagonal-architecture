package logger

import (
	"io"
)

type Config struct {
	name  string
	level *LogLevel

	writer io.Writer
}

type configFunc func(cfg *Config)

func WithLevel(level LogLevel) configFunc {
	return func(opts *Config) {
		opts.level = &level
	}
}

func WithName(name string) configFunc {
	return func(cfg *Config) {
		cfg.name = name
	}
}

func WithWriter(writer io.Writer) configFunc {
	return func(cfg *Config) {
		cfg.writer = writer
	}
}

func (cfg *Config) GetName() string {
	return cfg.name
}

func (cfg *Config) GetLevel() *LogLevel {
	return cfg.level
}

func (cfg *Config) GetWriter() io.Writer {
	return cfg.writer
}
