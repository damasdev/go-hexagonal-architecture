package user

import (
	"context"
	"fmt"
	port "go-hexagonal-architecture/internal/core/port/user"
	"go-hexagonal-architecture/internal/interfaces/constants"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	topic string
	conn  *amqp.Connection
}

func New(conn *amqp.Connection) port.UserConsumer {
	return &consumer{
		topic: constants.CONSUMER_TOPIC_USER,
		conn:  conn,
	}
}

func (c *consumer) Consume(ctx context.Context) {
	channel, err := c.conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		c.topic, // queue name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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
			fmt.Println("received messages:", string(response.Body))
		}
	}
}
