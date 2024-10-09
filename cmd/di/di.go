package di

import (
	"context"
	"fmt"
	"reflect"
	"yuemnoi-notification/db"
	"yuemnoi-notification/internal/config"
	"yuemnoi-notification/internal/event"
	"yuemnoi-notification/internal/handler"
	"yuemnoi-notification/internal/repository"
	"yuemnoi-notification/internal/route"
)

func must[T any](t T, err error) T {
	if err != nil {
		typeName := reflect.TypeOf(t).String()
		err := fmt.Errorf("failed to initialize %s: %w", typeName, err)
		panic(err)
	}
	return t
}

func InitDI(ctx context.Context, cfg *config.Config) (r *route.Handler, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// db
	dbPG := db.InitPostgreSQL(cfg)

	// repository
	userDeviceRepository := repository.NewUserDeviceRepository(dbPG)

	// handler
	userDeviceHandler := handler.NewUserDeviceHandler(userDeviceRepository)

	//event
	pushNotificationEvent := event.NewPushNotificationEvent(userDeviceRepository)

	r = route.NewHandler(userDeviceHandler)

	go func() {
		pushNotificationEvent.PushNotification(ctx, cfg)
	}()

	return r, nil
}
