package main

import (
	"os"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func main() {

	// init configuration
	config := config.New(".env")

	// init logger
	threshold := logger.LogLevel(config.GetInt("LOG_THRESHOLD"))
	logger := logger.New(os.Stdout, threshold)

	// setup and run server
	server.New(config, logger).Run()
}
