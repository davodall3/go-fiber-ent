package rabbitmq

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"projectSwagger/internal/model"
	"projectSwagger/internal/pkg/service"
)

type Consumer struct {
	RabbitMQ    RabbitMQ
	UserService service.UserService
}

func (c *Consumer) ListenAndServe() {
	queue, err := c.RabbitMQ.Channel.QueueDeclare(
		"user-queue",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Printf("Queue Declared About to %s\n", queue.Name)
	err = c.RabbitMQ.Channel.QueueBind(
		queue.Name,
		"users.event.*",
		"users",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	fmt.Printf("Queue Binding About to %s\n", queue.Name)

	msgs, err := c.RabbitMQ.Channel.Consume(
		queue.Name,
		"user-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	fmt.Printf("Consumer Registred About to queue %s\n", queue.Name)

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a RoutingKey: %s", msg.RoutingKey)
			switch msg.RoutingKey {
			case "users.event.create":
				user, err := decodeUserBody(msg.Body)
				if err != nil {
					log.Println(err)
				}
				fmt.Println("Created user", user)
			}
		}

		<-forever
	}()

}

func decodeUserBody(body []byte) (*model.UserBody, error) {
	var usr *model.UserBody
	if err := gob.NewDecoder(bytes.NewReader(body)).Decode(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}
