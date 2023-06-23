package main

import (
	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func init() {
	config.LoadEnvVars()

	logger.Initialize(
		logger.WithLevel(logger.InfoLevel),
		logger.WithName("fiber"),
	)
}

func main() {
	logger.Logger.Info("p")
	server.NewFiber().Run()
}
