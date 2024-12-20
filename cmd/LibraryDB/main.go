package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres:///library")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()
}
