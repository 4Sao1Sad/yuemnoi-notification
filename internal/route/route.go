package route

import (
	"yuemnoi-notification/internal/config"
	"yuemnoi-notification/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userDeviceHandler *handler.UserDeviceHandler
}

func NewHandler(userDeviceHandler *handler.UserDeviceHandler) *Handler {
	return &Handler{
		userDeviceHandler: userDeviceHandler,
	}
}

func (h *Handler) RegisterRouter(r fiber.Router, cfg *config.Config) {
	{
		r.Post("/user-device", h.userDeviceHandler.CreateUserDevice)
	}
}
