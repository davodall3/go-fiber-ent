package controller

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/app/model"
)

// CreateUser Creating User
//
//	@Summary		Creating User
//	@Description	Creating User with given request
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CreateUserResponse	true	"Request of Creating User Object"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/users [post]
func CreateUser(app *fiber.App) fiber.Router {
	return app.Post("/users", func(ctx *fiber.Ctx) error {
		var request model.CreateUserResponse
		err := ctx.BodyParser(&request)
		if err != nil {
			return err
		}
		//client := client.Client{}

		return ctx.Status(fiber.StatusCreated).JSON("User created successfully!")
	})
}
