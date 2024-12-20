package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetBooks(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT book_id, title, publication_year, genre FROM book")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching books: %v", err)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.PublicationYear, &book.Genre); err != nil {
			c.String(http.StatusInternalServerError, "Error scanning book: %v", err)
			return
		}
		books = append(books, book)
	}

	if rows.Err() != nil {
		c.String(http.StatusInternalServerError, "Error iterating over rows: %v", rows.Err())
		return
	}

	c.HTML(http.StatusOK, "books.html", gin.H{"Books": books})
}
