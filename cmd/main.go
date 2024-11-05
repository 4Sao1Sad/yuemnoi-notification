package main

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"yuemnoi-notification/cmd/di"
	"yuemnoi-notification/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recover from panic!",
				slog.Any("err", r),
			)
		}
	}()

	cfg := config.Load()
	app := fiber.New()
	handler, err := di.InitDI(ctx, cfg)
	if err != nil {
		slog.Error("failed to initialize DI, exiting...",
			"error", err)
		os.Exit(1)
		return
	}
	app.Use(requestid.New())

	handler.RegisterRouter(app, cfg)

	app.Listen(":" + strconv.Itoa(int(cfg.Port)))
}
