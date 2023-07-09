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
	opts := logger.Configs{}

	// Test initial values
	assert.Nil(t, opts.GetName())
	assert.Nil(t, opts.GetLevel())
	assert.Nil(t, opts.GetWriter())

	// Test WithLevel function
	level := logger.InfoLevel
	logger.WithLevel(level)(&opts)
	assert.Equal(t, level, *opts.GetLevel())

	// Test WithName function
	name := "mylogger"
	logger.WithName(name)(&opts)
	assert.Equal(t, name, *opts.GetName())

	// Test WithWriter function
	writer := mocks.NewMockWriter(bytes.NewBuffer(nil))
	logger.WithWriter(writer)(&opts)
	assert.Equal(t, writer, opts.GetWriter())
}
