package api

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start server",
		Action: func(c *cli.Context) error {
			startServer()
			return nil
		},
	}
}
