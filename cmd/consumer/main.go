package main

import (
	"context"
	"fmt"
	userConsumer "go-hexagonal-architecture/internal/infrastructure/consumer/rabbitmq"
	"go-hexagonal-architecture/pkg/config"
	"go-hexagonal-architecture/pkg/constants"
)

func init() {
	config.LoadEnvVars()
	config.ConnectRabbit()
}

func main() {
	conn := config.RabbitConn
	defer conn.Close()

	ctx := context.Background()

	userConsumer := userConsumer.New(conn)

	userConsumer.Subscribe(ctx, constants.CONSUMER_TOPIC_USER, func(message []byte) error {
		fmt.Println("received message:", string(message))
		return nil
	})

	<-ctx.Done()
}
