package main

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"yuemnoi-notification/dto"
	"yuemnoi-notification/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.Load()
	notification := dto.NotificationRequest{
		Message: "hello notification",
		UserIds: []int{1},
	}

	// Marshal the notification struct into JSON
	body, err := json.Marshal(notification)
	if err != nil {
		log.Fatalf("[client]: failed to marshal notification %+v", err)
	}

	conn, err := amqp.Dial(cfg.RabbitMQUrl)
	if err != nil {
		log.Fatalf("[client]: unable to connect RabbitMQ %+v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[client]: failed to open a channel %+v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("[client]: failed to declare a queue %+v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("[client]: failed to publish a message %+v", err)
	}

	log.Print(" [x] Notification Sent \n")
}
