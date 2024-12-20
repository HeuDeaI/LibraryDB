package main

import (
	"context"
	"log"
	"net/http"

	"LibraryDB/internal"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://heudeai:@localhost:5432/library?sslmode=disable"
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Library Management System")
	})
	r.GET("/books", func(c *gin.Context) { internal.GetBooks(c, dbPool) })
	r.Run(":8080")
}
