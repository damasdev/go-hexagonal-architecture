package api

import (
	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/server"
)

func init() {
	config.LoadEnvVars()
}

func Run() {
	server.NewFiber().Run()
}
