package handler

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/app/model"
	"projectSwagger/internal/app/pkg/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

// LoginUserHandler LoginUser login user
//
//	@Summary		Login User
//	@Description	Login Users
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.LoginUserRequest	true	"Request of Creating User Object"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/login [post]
func (h *AuthHandler) LoginUserHandler(c *fiber.Ctx) error {
	payload := new(model.LoginUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't login",
			"error":   err.Error(),
		})
	}

	// check username
	newUser, err := h.AuthService.Login(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": newUser,
	})
}
