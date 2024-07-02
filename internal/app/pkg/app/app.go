package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"projectSwagger/internal/app/controller"
	"projectSwagger/internal/app/mw"
	"projectSwagger/internal/app/pkg/client"
)

type App struct {
	Fiber *fiber.App
}

func New() (*App, error) {
	app := App{}
	app.Fiber = fiber.New()
	app.Fiber.Get("/swagger/*", swagger.HandlerDefault)

	app.Fiber.Use(recover.New())
	app.Fiber.Use(cors.New())

	// mw
	mw.AddCorrelationId(app.Fiber)

	// endpoints
	controller.CreateUser(app.Fiber)

	err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	return &app, nil
}

func (a *App) Run() error {
	fmt.Println("server running on :8080")
	err := a.Fiber.Listen(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
