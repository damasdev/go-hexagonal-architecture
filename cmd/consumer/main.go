package main

import (
	userConsumer "go-hexagonal-architecture/internal/interfaces/pubsub/user"
	"go-hexagonal-architecture/pkg/config"
	"log"
)

func init() {
	config.LoadEnvVars()
	config.ConnectRabbit()
}

func main() {
	userConsumer := userConsumer.NewConsumer(config.RabbitConn)

	err := userConsumer.StartConsumer(1, "user-exchange", "user-queue", "user-key", "user-tag")
	if err != nil {
		log.Fatal(err.Error())
	}
}
