package rabbitmq

import (
	"bytes"
	"encoding/gob"
	"log"
	"projectSwagger/internal/app/model"
	"projectSwagger/internal/app/pkg/service"
)

type Consumer struct {
	RabbitMQ    *RabbitMQ
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
	err = c.RabbitMQ.Channel.QueueBind(
		queue.Name,
		"users.event.*",
		"users",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

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

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.RoutingKey)
			switch msg.RoutingKey {
			case "users.event.create":
				user, err := decodeUserBody(msg.Body)
				if err != nil {
					log.Println(err)
				}
				_, err = c.UserService.CreateUser(user)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

}

func decodeUserBody(body []byte) (*model.UserBody, error) {
	var usr *model.UserBody
	if err := gob.NewDecoder(bytes.NewReader(body)).Decode(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}
