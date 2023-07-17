package main

import (
	"context"
	"go-hexagonal-architecture/internal/interfaces/pubsub/consumer"
	userHandler "go-hexagonal-architecture/internal/interfaces/pubsub/handler/user"
	"go-hexagonal-architecture/pkg/config"

	"log"
)

func init() {
	config.LoadEnvVars()
	config.ConnectRabbit()
}

func main() {

	userHandler := userHandler.New()

	consumer := consumer.New(config.RabbitConn)

	consumer.RegisterHandler("user-exchange", "user-queue-hello", "user-key-hello", "user-tag-hello", userHandler.Hello)
	consumer.RegisterHandler("user-exchange", "user-queue-world", "user-key-world", "user-tag-world", userHandler.World)

	if err := consumer.Start(context.Background()); err != nil {
		log.Fatal(err.Error())
	}
}
