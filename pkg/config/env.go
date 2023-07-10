package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// If you call Load without any args it will default to loading .env in the current path.
func LoadEnvVars(filenames ...string) {

	fileName := "./env"
	if len(filenames) > 0 {
		fileName = filenames[0]
	}

	// Get the current working directory
	dir, _ := os.Getwd()

	// Set the appPath to the current working directory
	appPath := filepath.Join(dir, fileName)

	godotenv.Load(appPath)
}
