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
func LoanBook(c *gin.Context, dbPool *pgxpool.Pool) {
	readerData := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone_number"`
		Email     string `json:"email"`
		BookID    string `json:"book_id"`
	}{}

	if err := c.ShouldBindJSON(&readerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tx, err := dbPool.Begin(c)
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer tx.Rollback(c)

	var readerID int
	err = tx.QueryRow(c, `
        SELECT reader_id FROM Reader 
        WHERE first_name=$1 AND last_name=$2 AND email=$3`,
		readerData.FirstName, readerData.LastName, readerData.Email,
	).Scan(&readerID)

	if err != nil { // Reader doesn't exist; insert them.
		err = tx.QueryRow(c, `
            INSERT INTO Reader (first_name, last_name, phone_number, email)
            VALUES ($1, $2, $3, $4) RETURNING reader_id`,
			readerData.FirstName, readerData.LastName, readerData.Phone, readerData.Email,
		).Scan(&readerID)

		if err != nil {
			log.Printf("Error adding reader: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	_, err = tx.Exec(c, `
        INSERT INTO Loan (book_id, reader_id, issue_date, return_date)
        VALUES ($1, $2, CURRENT_DATE, CURRENT_DATE + INTERVAL '7 days')`,
		readerData.BookID, readerID,
	)
	if err != nil {
		log.Printf("Error adding loan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err = tx.Commit(c); err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book loaned successfully"})
}
