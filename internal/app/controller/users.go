package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"projectSwagger/internal/app/model"
)

func GetUsers(app *fiber.App) fiber.Router {
	return app.Get("/users/all", func(ctx *fiber.Ctx) error {
		db, err := sql.Open("sqlite3", "postgres://user:pass@localhost/bookstore")
		if err != nil {
			log.Fatal(err)
		}
		var users []model.UserBody
		rows, err := db.Query("SELECT * FROM users")
		defer rows.Close()
		for rows.Next() {
			user := new(model.UserBody)
			users = append(users, *user)
			log.Println("aaaaaaa", user)
		}
		return ctx.Status(fiber.StatusCreated).JSON("User created successfully!")
	})
}

func CreateUser(app *fiber.App) fiber.Router {
	return app.Post("/users", func(ctx *fiber.Ctx) error {
		var request model.UserBody
		err := ctx.BodyParser(&request)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusCreated).JSON("User created successfully!")
	})
}
