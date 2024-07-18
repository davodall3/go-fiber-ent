package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "projectSwagger/docs"
	"projectSwagger/internal/app/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
