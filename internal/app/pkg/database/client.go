package database

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"projectSwagger/ent"
	"time"
)

const (
	username = "docker"
	password = "Milimada1956!"
	hostname = "godockerDB:3306"
	dbname   = "godocker"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func DBCreation() (*ent.Client, error) {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, nil
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)

	client := ent.NewClient(ent.Driver(drv))

	fmt.Println("db created")
	//// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
