package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitConn *amqp.Connection

type Rabbit struct {
	Host     string
	User     string
	Password string
}

func LoadRabbitConfig() Rabbit {
	conf := Rabbit{
		Host:     os.Getenv("RABBIT_HOST"),
		User:     os.Getenv("RABBIT_USERNAME"),
		Password: os.Getenv("RABBIT_PASSWORD"),
	}
	return conf
}

func ConnectRabbit() {

	cfg := LoadRabbitConfig()
	var err error

	connString := "amqps://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host

	// Connect to RabbitMQ
	RabbitConn, err = amqp.Dial(connString)
	if err != nil {
		log.Fatal(err)
	}
}
