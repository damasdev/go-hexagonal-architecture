package main

import (
	"os"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func init() {
	config.LoadEnvVars()
	logger.New(os.Stdout, logger.InfoLevel)
}

func main() {
	server.New().Run()
}
