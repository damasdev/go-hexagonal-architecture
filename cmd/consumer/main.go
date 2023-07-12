package main

import (
	"context"
	userConsumer "go-hexagonal-architecture/internal/infrastructure/consumer/rabbitmq"
	"go-hexagonal-architecture/pkg/config"
)

func init() {
	config.LoadEnvVars()
	config.ConnectRabbit()
}

func main() {
	// Instance
	conn := config.RabbitConn
	defer conn.Close()

	ctx := context.Background()

	// Consumer
	userConsumer := userConsumer.New(conn)
	userConsumer.Consume(ctx)

	<-ctx.Done()
}
