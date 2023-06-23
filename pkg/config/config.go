package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type config struct{}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal(err)
	}

	return &config{}
}

func (cfg *config) GetString(key string) string {
	return os.Getenv(key)
}

func (cfg *config) GetInt(key string) int {
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return 0
	}

	return int(value)
}

func (cfg *config) GetBool(key string) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return false
	}

	return value
}
