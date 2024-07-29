package rabbitmq

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"projectSwagger/internal/app/model"
	"time"
)

// RabbitMQ ...
type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQ ...
func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	fmt.Println("Successfully connected to RabbitMQ")
	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	fmt.Println("Successfully open a channel")

	err = channel.ExchangeDeclare(
		"users",  // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to Exchange Declare")
	fmt.Println("Successfully Exchange Declare")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")
	fmt.Println("Successfully Set QoS")

	return &RabbitMQ{
		Connection: conn,
		Channel:    channel,
	}
}

func (p *RabbitMQ) CreateUser(body *model.UserBody) error {
	fmt.Println("User created RabbitMQ")
	return p.publish("users.event.create", body)
}

func (p *RabbitMQ) publish(routingKey string, event interface{}) error {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.Channel.Publish(
		"users",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        eventJson,
			Timestamp:   time.Now(),
		})
	if err != nil {
		return err
	}
	return nil
}
