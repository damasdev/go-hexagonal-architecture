package main

import (
	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/server"
)

func init() {
	config.LoadEnvVars()
}

func main() {
	server.NewFiber().Run()
}
