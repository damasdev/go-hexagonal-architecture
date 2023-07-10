package logger_test

import (
	"bytes"
	"testing"

	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConfigs(t *testing.T) {
	// Create a new instance of configs
	cfg := logger.Config{}

	// Test initial values
	assert.Empty(t, cfg.GetName())
	assert.Nil(t, cfg.GetLevel())
	assert.Nil(t, cfg.GetWriter())

	// Test WithLevel function
	level := logger.InfoLevel
	logger.WithLevel(level)(&cfg)
	assert.Equal(t, level, *cfg.GetLevel())

	// Test WithName function
	name := "mylogger"
	logger.WithName(name)(&cfg)
	assert.Equal(t, name, cfg.GetName())

	// Test WithWriter function
	writer := mocks.NewMockWriter(bytes.NewBuffer(nil))
	logger.WithWriter(writer)(&cfg)
	assert.Equal(t, writer, cfg.GetWriter())

	// Test WithSkip function
	skip := 1
	logger.WithSkip(skip)(&cfg)
	assert.Equal(t, skip, cfg.GetSkip())
}
