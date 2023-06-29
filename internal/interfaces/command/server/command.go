package server

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Command for starting server",
		Action: func(ctx *cli.Context) error {
			return runServer()
		},
	}
}
