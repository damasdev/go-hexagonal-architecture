package user

import (
	"context"
	port "go-hexagonal-architecture/internal/core/port/user"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	conn *amqp.Connection
}

func New(conn *amqp.Connection) port.UserConsumer {
	return &consumer{
		conn: conn,
	}
}

func (c *consumer) Subscribe(ctx context.Context, name string, handler func(message []byte) error) {
	channel, err := c.conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		name,  // queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	delivery, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatal(err)
	}

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case response := <-delivery:
			err := handler(response.Body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
