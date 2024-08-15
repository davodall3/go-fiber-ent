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
	"sync"
	"syscall"
)

func SetupRoutes(app *fiber.App) {
	var wg sync.WaitGroup
	// DB database
	db, _ := database.DBCreation()

	// Add mock data
	database.AddProducts(db)

	// Services
	userService := service.NewUserService(db)
	authService := service.NewAuthService(db)
	productService := service.NewProductService(db)

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
	go func() {
		wg.Add(1)
		app.Post("/products/buy", productHandler.BuyProductHandler)
		wg.Done()
	}()

	wg.Wait()

	errC := make(chan error, 1)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()
		defer func() {
			fmt.Println("Closing all connections")
			rabbitMQ.Channel.Close()
			rabbitMQ.Connection.Close()
			db.Close()
			stop()
			close(errC)
		}()
	}()
}
