package user

import (
	"errors"

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
)

type consumer struct {
	conn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) port.UserConsumer {
	return &consumer{
		conn: conn,
	}
}

func (c *consumer) StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error {
	return nil
}

func (c *consumer) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, errors.New("error consumer.createchannel on conn.channel")
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
		return nil, errors.New("error consumer.createchannel on ch.exchangedeclare")
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
		return nil, errors.New("error consumer.createchannel on ch.queuedeclare")
	}

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.createchannel on ch.queuebind")
	}

	err = ch.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.New("error consumer.createchannel on ch.Qos")
	}

	return ch, nil
}
