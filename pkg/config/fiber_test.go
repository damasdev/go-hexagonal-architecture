package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFiberConfig(t *testing.T) {
	// Set the required environment variables
	os.Setenv("APP_NAME", "MyApp")
	os.Setenv("SERVER_READ_TIMEOUT", "10")
	os.Setenv("SERVER_WRITE_TIMEOUT", "20")
	os.Setenv("SERVER_IDLE_TIMEOUT", "30")
	os.Setenv("SERVER_PREFORK", "true")

	// Call the FiberConfig function
	config := config.FiberConfig()

	// Assert the configuration values
	assert.Equal(t, "MyApp", config.AppName)
	assert.Equal(t, time.Second*10, config.ReadTimeout)
	assert.Equal(t, time.Second*20, config.WriteTimeout)
	assert.Equal(t, time.Second*30, config.IdleTimeout)
	assert.True(t, config.Prefork)

	// Clean up the environment variables
	os.Unsetenv("APP_NAME")
	os.Unsetenv("SERVER_READ_TIMEOUT")
	os.Unsetenv("SERVER_WRITE_TIMEOUT")
	os.Unsetenv("SERVER_IDLE_TIMEOUT")
	os.Unsetenv("SERVER_PREFORK")
}
