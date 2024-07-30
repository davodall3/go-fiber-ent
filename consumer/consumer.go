package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"projectSwagger/internal/app/model"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	fmt.Println("Successfully connected to RabbitMQ")
	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	fmt.Println("Successfully open a channel")

	queue, err := channel.QueueDeclare(
		"user-queue",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Printf("Queue Declared About to %s\n", queue.Name)

	err = channel.QueueBind(
		queue.Name,
		"users.event.*",
		"users",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	fmt.Printf("Queue Binding About to %s\n", queue.Name)

	msgs, err := channel.Consume(
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

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" Body == %s", d.Body)
			log.Printf(" RoutingKey == %s", d.RoutingKey)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func decodeUserBody(body []byte) (*model.UserBody, error) {
	var usr *model.UserBody
	if err := gob.NewDecoder(bytes.NewReader(body)).Decode(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
