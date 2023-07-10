package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// If you call Load without any args it will default to loading .env in the current path.
func LoadEnvVars(filenames ...string) {

	// Set file name
	fileName := "/.env"
	if len(filenames) > 0 {
		fileName = filenames[0]
	}

	// Get the current working directory
	dir, _ := os.Getwd()

	// Set the appPath to the current working directory
	appPath := filepath.Join(dir, fileName)

	// Check file
	if _, err := os.Stat(appPath); err != nil {
		log.Fatal(err)
	}

	// Load env
	godotenv.Load(appPath)
}
