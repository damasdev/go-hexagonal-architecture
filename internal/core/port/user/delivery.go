package user

type UserPublisher interface {
	Publish(body []byte, contentType string) error
}

type UserConsumer interface {
	StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error
}
