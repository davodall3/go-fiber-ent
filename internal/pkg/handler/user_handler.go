package handler

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/model"
	"projectSwagger/internal/pkg/rabbitmq"
	"projectSwagger/internal/pkg/service"
)

type UserHandler struct {
	UserService service.UserService
	Producer    rabbitmq.RabbitMQ
}

func NewUserHandler(service service.UserService, producer rabbitmq.RabbitMQ) *UserHandler {
	return &UserHandler{
		UserService: service,
		Producer:    producer,
	}
}

// CreateUserHandler Creating User
//
//	@Summary		Creating User
//	@Description	Creating User with given request
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserBody	true	"Request of Creating User Object"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/users [post]
func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	payload := new(model.UserBody)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Payload error",
			"error":   err.Error(),
		})
	}

	user, err := h.UserService.CreateUser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
			"error":   err.Error(),
		})
	}

	err = h.Producer.CreateUser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": user,
	})
}

// GetAllUsersHandler GetUser getting user
//
//	@Summary		Getting Users
//	@Description	Getting Users with given request
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		400	{string}	string	"Bad Request"
//	@Router			/users/all [get]
func (h *UserHandler) GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.Producer.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}
