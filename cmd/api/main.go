package main

import (
	"os"
	"strconv"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func main() {

	// init configuration
	config := config.New()

	// setup logger threshold
	level, err := strconv.ParseUint(config.Get("LOG_LEVEL"), 10, 8)
	if err != nil {
		level = uint64(logger.WarnLevel)
	}

	// init logger
	logger := logger.New(os.Stdout, logger.LogLevel(level))

	// setup and run server
	server.New(config, logger).Run()
}
