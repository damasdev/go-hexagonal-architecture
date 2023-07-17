package user

import (
	"context"
	"errors"
	"fmt"

	port "go-hexagonal-architecture/internal/core/port/user"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeKind       = "direct"
	exchangeDurable    = true
	exchangeAutoDelete = false
	exchangeInternal   = false
	exchangeNoWait     = false

	queueDurable    = true
	queueAutoDelete = false
	queueExclusive  = false
	queueNoWait     = false

	prefetchCount  = 1
	prefetchSize   = 0
	prefetchGlobal = false

	consumeAutoAck   = false
	consumeExclusive = false
	consumeNoLocal   = false
	consumeNoWait    = false
)

type consumer struct {
	amqpConn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) port.UserConsumer {
	return &consumer{
		amqpConn: conn,
	}
}

func (c *consumer) StartConsumer(workerPoolSize int, exchangeName, queueName, routingKey, consumerTag string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := c.CreateChannel(exchangeName, queueName, routingKey)
	if err != nil {
		return errors.New("error consumer.StartConsumer on c.CreateChannel")
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		consumeAutoAck,
		consumeExclusive,
		consumeNoLocal,
		consumeNoWait,
		nil,
	)
	if err != nil {
		return errors.New("error consumer.StartConsumer on ch.Consume")
	}

	for i := 0; i < workerPoolSize; i++ {
		go c.worker(ctx, deliveries)
	}

	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}

func (c *consumer) CreateChannel(exchangeName, queueName, routingKey string) (*amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on amqpConn.Channel")
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.ExchangeDeclare")
	}

	queue, err := ch.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.QueueDeclare")
	}

	err = ch.QueueBind(
		queue.Name,
		routingKey,
		exchangeName,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.QueueBind")
	}

	err = ch.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.Qos")
	}

	return ch, nil
}

func (c *consumer) worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		fmt.Println("received message:", string(delivery.Body))
		delivery.Ack(false)
	}
}
