package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "X-Total-Count",
	}))

	app.Fiber.Static("/", "./admin/dist", fiber.Static{
		Index: "index.html",
	})

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
