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
	dbURL := "postgres:///library"
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/books-with-authors", func(c *gin.Context) { internal.GetBooksWithAuthors(c, dbPool) })
	r.GET("/readers", func(c *gin.Context) { internal.GetReaders(c, dbPool) })
	r.GET("/loans", func(c *gin.Context) { internal.GetLoans(c, dbPool) })
	r.GET("/book/:bookID", func(c *gin.Context) {
		c.HTML(http.StatusOK, "book.html", nil)
	})
	r.GET("/book-data/:bookID", func(c *gin.Context) { internal.GetBook(c, dbPool) })
	r.POST("/loan-book", func(c *gin.Context) { internal.LoanBook(c, dbPool) })
	r.GET("/add-book", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_book.html", nil)
	})
	r.POST("/add-book", func(c *gin.Context) { internal.AddBook(c, dbPool) })

	r.Run(":8080")
}
