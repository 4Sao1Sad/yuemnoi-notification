package event

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func InitFirebaseClient(ctx context.Context) (*messaging.Client, error) {
	// Path to the service account key JSON file
	filePath := "internal/event/secret/sa-notification-8d649-firebase-adminsdk-jucjc-48a166bc1f.json"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("error file not found: %v\n", err)
	}

	opt := option.WithCredentialsFile(filePath)
	fmt.Println("opt", opt)

	// Initialize Firebase App
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	// Get Firebase messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting messaging client: %v\n", err)
		return nil, err
	}

	return client, nil
}
