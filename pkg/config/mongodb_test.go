package config_test

import (
	"os"
	"testing"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadMongoConfig(t *testing.T) {
	// Set the required environment variables
	os.Setenv("MONGO_DRIVER", "mongodb")
	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_USERNAME", "admin")
	os.Setenv("MONGO_PASSWORD", "password")
	os.Setenv("MONGO_DBNAME", "mydb")
	os.Setenv("MONGO_SRV", "false")

	// Call the LoadMongoConfig function
	conf := config.LoadMongoConfig()

	// Assert the config values
	assert.Equal(t, "mongodb", conf.Driver)
	assert.Equal(t, "localhost", conf.Host)
	assert.Equal(t, "27017", conf.Port)
	assert.Equal(t, "admin", conf.User)
	assert.Equal(t, "password", conf.Password)
	assert.Equal(t, "mydb", conf.DBName)
	assert.Equal(t, "false", conf.UseSRV)

	// Clean up the environment variables
	os.Unsetenv("MONGO_DRIVER")
	os.Unsetenv("MONGO_HOST")
	os.Unsetenv("MONGO_PORT")
	os.Unsetenv("MONGO_USERNAME")
	os.Unsetenv("MONGO_PASSWORD")
	os.Unsetenv("MONGO_DBNAME")
	os.Unsetenv("MONGO_SRV")
}

func TestConnectMongoDB(t *testing.T) {
	// load environment
	config.LoadEnvVars("./../../.env")

	// Call the ConnectMongoDB function
	config.ConnectMongoDB()

	// Assert that the MongoDB variable is set and not nil
	assert.NotNil(t, config.MongoDB)
}
