package main

import (
	"LibraryDB/internal"
	"context"
	"log"
)

func main() {
	// Establish connection to the database
	conn, err := internal.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
		return
	}

	// Ensure that the connection is closed after all operations are complete
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()
}
