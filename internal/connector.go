package internal

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Connect establishes a connection to the PostgreSQL database using the DATABASE_URL
// from the environment variables loaded by godotenv.
// It returns a pgx.Conn object and an error if any issues occur during connection.
func Connect() (*pgx.Conn, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v\n", err) // Use log for better visibility
	}

	// Get the DATABASE_URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set in environment variables")
	}

	// Attempt to connect to the PostgreSQL database
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database.")
	return conn, nil
}
