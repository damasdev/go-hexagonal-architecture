package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type config struct{}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func New() Config {
	return &config{}
}

func (cfg *config) Get(key string) string {
	return os.Getenv(key)
}
