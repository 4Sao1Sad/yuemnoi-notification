package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"yuemnoi-notification/dto"
	"yuemnoi-notification/internal/config"
	"yuemnoi-notification/internal/repository"

	"firebase.google.com/go/messaging"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PushNotificationEvent struct {
	UserDeviceRepository repository.UserDeviceRepository
}

func NewPushNotificationEvent(UserDeviceRepository repository.UserDeviceRepository) *PushNotificationEvent {
	return &PushNotificationEvent{UserDeviceRepository}
}

func (h PushNotificationEvent) PushNotification(ctx context.Context, cfg *config.Config) {
	queueName := "notification"

	conn, err := amqp.Dial(cfg.RabbitMQUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	client, err := InitFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("failed to init firebase client: %v\n", err)
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			var notification dto.NotificationRequest
			if err := json.Unmarshal(msg.Body, &notification); err != nil {
				log.Printf("Error unmarshalling message: %s", err)
				continue
			}

			log.Printf("Received a message: %s", notification.Message)

			fmt.Println("userIds", notification.UserIds)
			for _, userId := range notification.UserIds {
				tokens, err := h.UserDeviceRepository.GetUserDevices(userId)
				if err != nil {
					log.Printf("Error get device token: %s", err)
					continue
				}

				fmt.Println("tokens", tokens)
				for _, token := range tokens {
					message := &messaging.Message{
						Token: token.DeviceToken,
						Notification: &messaging.Notification{
							Title: "Hello11111111",
							Body:  "This is a push notification",
						},
					}

					// Send the message
					response, err := client.Send(context.Background(), message)
					if err != nil {
						log.Fatalf("error sending message: %v\n", err)
					}

					// Response is a message ID string
					log.Printf("Successfully sent message: %s\n", response)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
