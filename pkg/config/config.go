package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	dir, _ := os.Getwd()
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
}
