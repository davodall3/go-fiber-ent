package router

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"projectSwagger/internal/pkg/database"
	"projectSwagger/internal/pkg/handler"
	"projectSwagger/internal/pkg/rabbitmq"
	"projectSwagger/internal/pkg/service"
	"syscall"
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

	// RabbitMQ
	rabbitMQ := rabbitmq.NewRabbitMQ()

	// Handlers
	userHandler := handler.NewUserHandler(*userService, *rabbitMQ)
	authHandler := handler.NewAuthHandler(*authService)
	productHandler := handler.NewProductHandler(*productService)

	// API
	app.Post("/users", userHandler.CreateUserHandler)
	app.Get("/users/all", userHandler.GetAllUsersHandler)
	app.Post("/login", authHandler.LoginUserHandler)
	app.Get("/products/all", productHandler.GetAllProductsHandler)
	app.Post("/products/buy", productHandler.BuyProductHandler)

	errC := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		defer func() {
			fmt.Println("Closing all connections")
			rabbitMQ.Channel.Close()
			rabbitMQ.Connection.Close()
			client.Close()
			stop()
			close(errC)
		}()
	}()
}
