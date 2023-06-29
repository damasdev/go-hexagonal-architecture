package main

import (
	"log"

	"github.com/damasdev/fiber/internal/interfaces/command/server"
	"github.com/damasdev/fiber/pkg/cli"
	"github.com/damasdev/fiber/pkg/config"
)

func init() {
	config.LoadEnvVars()
}

func main() {

	cli := cli.NewCLI()

	// register all command-line
	cli.RegisterCommand(
		server.Command(),
	)

	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
