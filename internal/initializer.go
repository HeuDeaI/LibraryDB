package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

const (
	tablesSQLFile = "initTables.sql"
	dataSQLFile   = "initData.sql"
)

// InitializeTables creates tables and inserts data into the database
func InitializeTables(conn *pgx.Conn) error {
	// Get file paths for SQL files
	tablesSQLPath, err := getFilePath(tablesSQLFile)
	if err != nil {
		return fmt.Errorf("unable to get path for tables SQL file: %w", err)
	}

	dataSQLPath, err := getFilePath(dataSQLFile)
	if err != nil {
		return fmt.Errorf("unable to get path for data SQL file: %w", err)
	}

	// Initialize tables
	if err := executeSQLFile(conn, tablesSQLPath); err != nil {
		return fmt.Errorf("error initializing tables: %w", err)
	}

	// Insert data
	if err := executeSQLFile(conn, dataSQLPath); err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}

	return nil
}

// executeSQLFile reads and executes an SQL file
func executeSQLFile(conn *pgx.Conn, filePath string) error {
	sqlContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read SQL file %s: %w", filePath, err)
	}

	_, err = conn.Exec(context.Background(), string(sqlContent))
	if err != nil {
		return fmt.Errorf("error executing SQL file %s: %w", filePath, err)
	}

	log.Printf("Executed SQL file: %s\n", filePath)
	return nil
}

// getFilePath constructs the absolute path for SQL files
func getFilePath(fileName string) (string, error) {
	// Build the absolute path for the SQL files in the internal directory
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get current working directory: %w", err)
	}
	return filepath.Join(dir, "internal", fileName), nil
}
