package main

import (
	"LibraryDB/internal" // Adjust import path to match your project structure
	"context"
	"fmt"
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

	// Initialize tables and insert data
	if err := internal.InitializeTables(conn); err != nil {
		log.Fatalf("Failed to initialize tables and insert data: %v\n", err)
		return
	}

	// If initialization is successful
	fmt.Println("Database tables initialized and data inserted successfully!")
}
