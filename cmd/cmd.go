package cmd

import (
	"log"
	"os"

	"github.com/damasdev/fiber/cmd/api"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name: "app",
		Commands: []*cli.Command{
			api.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
