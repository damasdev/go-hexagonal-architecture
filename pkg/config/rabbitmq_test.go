package config_test

import (
	"os"
	"testing"

	"go-hexagonal-architecture/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestLoadRabbitConfig(t *testing.T) {
	// Set the required environment variables
	os.Setenv("RABBIT_HOST", "localhost")
	os.Setenv("RABBIT_USERNAME", "admin")
	os.Setenv("RABBIT_PASSWORD", "password")

	// Call the LoadRabbitConfig function
	conf := config.LoadRabbitConfig()

	// Assert the config values
	assert.Equal(t, "localhost", conf.Host)
	assert.Equal(t, "admin", conf.User)
	assert.Equal(t, "password", conf.Password)

	// Clean up the environment variables
	os.Unsetenv("RABBIT_HOST")
	os.Unsetenv("RABBIT_USERNAME")
	os.Unsetenv("RABBIT_PASSWORD")
}

func TestConnectRabbit(t *testing.T) {
	// load environment
	config.LoadEnvVars("./../../.env")

	// Call the ConnectRabbit function
	config.ConnectRabbit()

	// Assert that the RabbitConn variable is set and not nil
	assert.NotNil(t, config.RabbitConn)
}
