package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetBooksWithAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	query := `
		SELECT 
			b.book_id, b.title, b.publication_year, b.genre, 
			COALESCE(STRING_AGG(a.first_name || ' ' || a.last_name, ', '), 'Author unknown') AS authors
		FROM 
			book AS b
		LEFT JOIN 
			bookauthor AS ba ON b.book_id = ba.book_id
		LEFT JOIN 
			author AS a ON ba.author_id = a.author_id
		GROUP BY 
			b.book_id
		ORDER BY 
			b.book_id;
	`

	rows, err := dbPool.Query(c, query)
	if err != nil {
		log.Printf("Error fetching books with authors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var books []BookWithAuthors
	for rows.Next() {
		var book BookWithAuthors
		var authors string
		if err := rows.Scan(&book.BookID, &book.Title, &book.PublicationYear, &book.Genre, &authors); err != nil {
			log.Printf("Error scanning book with authors: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		book.Authors = append(book.Authors, authors)
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context, dbPool *pgxpool.Pool) {
	bookID := c.Param("bookID")
	query := `
        SELECT 
            b.book_id, b.title, b.publication_year, b.genre, 
            COALESCE(STRING_AGG(a.first_name || ' ' || a.last_name, ', '), 'Author unknown') AS authors
        FROM 
            book AS b
        LEFT JOIN 
            bookauthor AS ba ON b.book_id = ba.book_id
        LEFT JOIN 
            author AS a ON ba.author_id = a.author_id
        WHERE 
            b.book_id = $1
        GROUP BY 
            b.book_id;
    `
	var book BookWithAuthors
	var authors string
	err := dbPool.QueryRow(c, query, bookID).Scan(&book.BookID, &book.Title, &book.PublicationYear, &book.Genre, &authors)
	if err != nil {
		log.Printf("Error fetching book details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	book.Authors = append(book.Authors, authors)
	c.JSON(http.StatusOK, book)
}
