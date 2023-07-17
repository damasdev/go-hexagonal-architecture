package user

import (
	"context"
	"errors"
	"log"
	"time"

	port "go-hexagonal-architecture/internal/core/port/user"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	publishMandatory = false
	publishImmediate = false

	publishExchane    = "user-exchange"
	publishRoutingKey = "user-routing-key"
)

type publisher struct {
	exchangeName string
	routingKey   string

	amqpChan *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) port.UserPublisher {
	amqpChan, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &publisher{
		exchangeName: publishExchane,
		routingKey:   publishRoutingKey,

		amqpChan: amqpChan,
	}
}

func (p *publisher) Publish(ctx context.Context, body []byte, contentType string) error {
	if err := p.amqpChan.PublishWithContext(
		ctx,
		p.exchangeName,
		p.routingKey,
		publishMandatory,
		publishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         body,
		},
	); err != nil {
		return errors.New("error publisher.publish on ch.Publish")
	}

	return nil
}
