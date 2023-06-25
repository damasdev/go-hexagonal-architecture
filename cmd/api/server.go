package api

import "github.com/damasdev/fiber/pkg/server"

func Run() {
	server.NewFiber().Run()
}
