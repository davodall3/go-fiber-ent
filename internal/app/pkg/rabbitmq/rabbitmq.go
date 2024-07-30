package rabbitmq

import (
	"context"
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

func (r *RabbitMQ) CreateUser(body *model.UserBody) error {
	fmt.Println("User created RabbitMQ")
	return r.publish("users.event.create", body)
}

func (r *RabbitMQ) GetAllUser() error {
	fmt.Println("Get all users RabbitMQ")
	return r.publish("users.event.getAll", nil)
}

func (p *RabbitMQ) publish(routingKey string, event interface{}) error {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Println("json created ===>> ", string(eventJson))
	fmt.Println("routingKey ===>> ", routingKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err = p.Channel.PublishWithContext(ctx,
		"users",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        eventJson,
		})
	if err != nil {
		fmt.Println("publishing error === ", err)
		return err
	}
	fmt.Println("published ===>> ", string(eventJson))
	return nil
}
