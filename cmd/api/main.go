package main

import (
	"os"

	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
)

func main() {
	log := logger.New(os.Stdout, logger.InfoLevel)
	server.New(log).Run()
}
