package cli

import (
	"os"

	"github.com/urfave/cli/v2"
)

type Command interface {
	RegisterCommand(...*cli.Command)
	Run() error
}

type command struct {
	app *cli.App
}

func NewCLI() Command {
	return &command{
		app: &cli.App{
			Name:  os.Getenv("APP_NAME"),
			Usage: "Command-line interface for all service",
		},
	}
}

func (c *command) RegisterCommand(commands ...*cli.Command) {
	c.app.Commands = commands
}

func (c *command) Run() error {
	return c.app.Run(os.Args)
}
