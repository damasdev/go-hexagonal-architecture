package main

import "go-hexagonal-architecture/pkg/config"

func init() {
	config.LoadEnvVars()
	config.ConnectRabbit()
}

func main() {
	// consumer
}
