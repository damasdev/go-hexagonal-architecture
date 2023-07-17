package user

import (
	"log"

	port "go-hexagonal-architecture/internal/core/port/user"

	amqp "github.com/rabbitmq/amqp091-go"
)

type publisher struct {
	amqpChan *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) port.UserPublisher {
	amqpChan, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &publisher{
		amqpChan: amqpChan,
	}
}

func (p *publisher) Publish(body []byte, contentType string) error {
	return nil
}

func (p *publisher) CloseChan() {
	if err := p.amqpChan.Close(); err != nil {
		log.Fatalf("error publisher.closechan: %v", err)
	}
}
