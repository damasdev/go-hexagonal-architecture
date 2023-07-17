package consumer

import (
	"context"
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	RegisterHandler(exchange, queueName, bindingKey, consumerTag string, handler func(message []byte) error)

	Start(ctx context.Context) error
}

type Handler struct {
	handler func(message []byte) error
	config  struct {
		exchange    string
		queueName   string
		bindingKey  string
		consumerTag string
	}
}

type consumer struct {
	connection *amqp.Connection
	config     Config

	handlers map[string]Handler
}

func New(conn *amqp.Connection, config ...Config) Consumer {

	cfg := LoadDefaultConfig()

	if len(config) > 0 {
		cfg = config[0]
	}

	return &consumer{
		connection: conn,
		config:     cfg,
		handlers:   make(map[string]Handler),
	}
}

func (c *consumer) RegisterHandler(exchange, queueName, bindingKey, consumerTag string, handler func(message []byte) error) {
	c.handlers[consumerTag] = Handler{
		handler: handler,
		config: struct {
			exchange    string
			queueName   string
			bindingKey  string
			consumerTag string
		}{
			exchange:    exchange,
			queueName:   queueName,
			bindingKey:  bindingKey,
			consumerTag: consumerTag,
		},
	}
}

func (c *consumer) Start(ctx context.Context) error {

	for _, h := range c.handlers {
		go c.consumeMessages(ctx, h)
	}

	<-ctx.Done()

	return nil
}

func (c *consumer) consumeMessages(ctx context.Context, h Handler) {
	ch, err := c.CreateChannel(h.config.exchange, h.config.queueName, h.config.bindingKey)
	if err != nil {
		log.Fatal(errors.New("error consumer.StartConsumer on c.CreateChannel"))
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		h.config.queueName,
		h.config.consumerTag,
		c.config.consumeAutoAck,
		c.config.consumeExclusive,
		c.config.consumeNoLocal,
		c.config.consumeNoWait,
		nil,
	)
	if err != nil {
		log.Fatal(errors.New("error consumer.StartConsumer on ch.Consume"))
	}

	for delivery := range deliveries {

		err := h.handler(delivery.Body)
		if err != nil {
			log.Fatal(err)
		}

		delivery.Ack(false)
	}

	<-ch.NotifyClose(make(chan *amqp.Error))
}

func (c *consumer) CreateChannel(exchange, queueName, bindingKey string) (*amqp.Channel, error) {
	ch, err := c.connection.Channel()
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on connection.Channel")
	}

	err = ch.ExchangeDeclare(
		exchange,
		c.config.exchangeKind,
		c.config.exchangeDurable,
		c.config.exchangeAutoDelete,
		c.config.exchangeInternal,
		c.config.exchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.ExchangeDeclare")
	}

	queue, err := ch.QueueDeclare(
		queueName,
		c.config.queueDurable,
		c.config.queueAutoDelete,
		c.config.queueExclusive,
		c.config.queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.QueueDeclare")
	}

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchange,
		c.config.queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.QueueBind")
	}

	err = ch.Qos(
		c.config.prefetchCount,  // prefetch count
		c.config.prefetchSize,   // prefetch size
		c.config.prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.New("error consumer.CreateChannel on ch.Qos")
	}

	return ch, nil
}
