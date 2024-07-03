package router

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/app/pkg/database"
	"projectSwagger/internal/app/pkg/handler"
	"projectSwagger/internal/app/pkg/service"
)

func SetupRoutes(app *fiber.App) {
	// DB database
	client := database.DbConnection()

	// Services
	userService := service.NewUserService(client)
	authService := service.NewAuthService(client)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	authHandler := handler.NehAuthHandler(*authService)

	app.Post("/users", userHandler.CreateUserHandler)
	app.Get("/users/all", userHandler.GetAllUsersHandler)
	app.Post("/login", authHandler.LoginUserHandler)
}
