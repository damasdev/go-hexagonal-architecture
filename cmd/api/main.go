package main

import (
	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func init() {
	config.LoadEnvVars()

	logger.Initialize(
		logger.WithName("fiber"),
		logger.WithLevel(logger.InfoLevel),
	)
}

func main() {
	server.NewFiber().Run()
}
