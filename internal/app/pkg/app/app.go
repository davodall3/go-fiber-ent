package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"projectSwagger/internal/app/mw"
	"projectSwagger/internal/app/pkg/router"
)

type App struct {
	Fiber *fiber.App
}

func New() (*App, error) {
	app := App{}
	app.Fiber = fiber.New()
	app.Fiber.Get("/swagger/*", swagger.HandlerDefault)

	// MW
	mw.FiberMiddleware(app.Fiber)

	// Router
	router.SetupRoutes(app.Fiber)

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
