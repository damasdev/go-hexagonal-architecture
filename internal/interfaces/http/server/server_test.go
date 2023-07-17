package server_test

import (
	"go-hexagonal-architecture/internal/interfaces/http/server"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiberConfig(t *testing.T) {
	// Set the required environment variables
	os.Setenv("APP_NAME", "MyApp")
	os.Setenv("SERVER_PREFORK", "true")

	// Call the FiberConfig function
	config := server.LoadFiberConfig()

	// Assert the configuration values
	assert.Equal(t, "MyApp", config.AppName)
	assert.True(t, config.Prefork)

	// Clean up the environment variables
	os.Unsetenv("APP_NAME")
	os.Unsetenv("SERVER_PREFORK")
}
