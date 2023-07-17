package user

import "context"

type UserPublisher interface {
	Publish(ctx context.Context, body []byte, contentType string) error
}

type UserConsumer interface {
	StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error
}
