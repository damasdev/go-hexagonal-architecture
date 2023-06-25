package main

import (
	"github.com/damasdev/fiber/cmd"
	"github.com/damasdev/fiber/pkg/config"
)

func init() {
	config.LoadEnvVars()
}

func main() {
	cmd.Run()
}
