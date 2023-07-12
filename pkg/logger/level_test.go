package logger_test

import (
	"testing"

	"go-hexagonal-architecture/pkg/logger"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogLevel(t *testing.T) {
	// Test DebugLevel
	assert.Equal(t, zerolog.DebugLevel, logger.ToLevel(logger.DebugLevel))

	// Test InfoLevel
	assert.Equal(t, zerolog.InfoLevel, logger.ToLevel(logger.InfoLevel))

	// Test WarnLevel
	assert.Equal(t, zerolog.WarnLevel, logger.ToLevel(logger.WarnLevel))

	// Test ErrorLevel
	assert.Equal(t, zerolog.ErrorLevel, logger.ToLevel(logger.ErrorLevel))

	// Test PanicLevel
	assert.Equal(t, zerolog.PanicLevel, logger.ToLevel(logger.PanicLevel))

	// Test unknown LogLevel
	assert.Equal(t, zerolog.WarnLevel, logger.ToLevel(10))
}
