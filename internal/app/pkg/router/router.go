package router

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/app/pkg/database"
	"projectSwagger/internal/app/pkg/handler"
	"projectSwagger/internal/app/pkg/service"
)

func SetupRoutes(app *fiber.App) {
	// DB database
	client, _ := database.DBCreation()

	// Add mock data
	database.AddProducts(client)

	// Services
	userService := service.NewUserService(client)
	authService := service.NewAuthService(client)
	productService := service.NewProductService(client)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	authHandler := handler.NewAuthHandler(*authService)
	productHandler := handler.NewProductHandler(*productService)
	// API
	app.Post("/users", userHandler.CreateUserHandler)
	app.Get("/users/all", userHandler.GetAllUsersHandler)
	app.Post("/login", authHandler.LoginUserHandler)
	app.Get("/products/all", productHandler.GetAllProductsHandler)
	app.Post("/products/buy", productHandler.BuyProductHandler)
}
