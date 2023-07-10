package config_test

import (
	"os"
	"testing"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnvVars(t *testing.T) {
	config.LoadEnvVars("./../../.env")

	assert.NotEmpty(t, os.Getenv("APP_NAME"))
	assert.NotEmpty(t, os.Getenv("APP_PORT"))
}
