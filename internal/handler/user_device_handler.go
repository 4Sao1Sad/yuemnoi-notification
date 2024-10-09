package handler

import (
	"yuemnoi-notification/internal/model"
	"yuemnoi-notification/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserDeviceHandler struct {
	UserDeviceRepository repository.UserDeviceRepository
	Validator            *validator.Validate
}

func NewUserDeviceHandler(UserDeviceRepository repository.UserDeviceRepository) *UserDeviceHandler {
	return &UserDeviceHandler{
		UserDeviceRepository: UserDeviceRepository,
		Validator:            validator.New(),
	}
}

type UserDeviceRequest struct {
	UserId uint   `json:"userId" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

// CreateUserDevice handles creating a new user device
func (h UserDeviceHandler) CreateUserDevice(c *fiber.Ctx) error {
	var request UserDeviceRequest

	// Parse the JSON body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.Validator.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed: " + err.Error(),
		})
	}

	err := h.UserDeviceRepository.CreateUserDevice(model.UserDevice{
		UserID:      request.UserId,
		DeviceToken: request.Token,
	})

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// For demonstration, just returning a success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"userId": request.UserId,
		"token":  request.Token,
	})
}
