package config_test

import (
	"os"
	"testing"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFiberConfig(t *testing.T) {
	// Set the required environment variables
	os.Setenv("APP_NAME", "MyApp")
	os.Setenv("SERVER_PREFORK", "true")

	// Call the FiberConfig function
	config := config.FiberConfig()

	// Assert the configuration values
	assert.Equal(t, "MyApp", config.AppName)
	assert.True(t, config.Prefork)

	// Clean up the environment variables
	os.Unsetenv("APP_NAME")
	os.Unsetenv("SERVER_PREFORK")
}
