package internal

import (
	"bytes"
	"io"
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
	var readerData LoanRequest
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

func AddBook(c *gin.Context, dbPool *pgxpool.Pool) {
	var bookData BookData
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := c.ShouldBindJSON(&bookData); err != nil {
		log.Printf("Error binding JSON: %v", err)
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

	var authorIDs []int
	for _, author := range bookData.Authors {
		var authorID int
		err := tx.QueryRow(c, `
			INSERT INTO Author (first_name, last_name)
			VALUES ($1, $2)
			ON CONFLICT (first_name, last_name) DO UPDATE SET first_name = EXCLUDED.first_name
			RETURNING author_id`,
			author.FirstName, author.LastName,
		).Scan(&authorID)
		if err != nil { // Author doesn't exist; insert them.
			log.Printf("Error inserting author: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		authorIDs = append(authorIDs, authorID)
	}

	var bookID int
	err = tx.QueryRow(c, `
		INSERT INTO Book (title, publication_year, genre)
		VALUES ($1, $2, $3)
		RETURNING book_id`,
		bookData.Title, bookData.PublicationYear, bookData.Genre,
	).Scan(&bookID)
	if err != nil {
		log.Printf("Error inserting book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	for _, authorID := range authorIDs {
		_, err := tx.Exec(c, `
			INSERT INTO BookAuthor (book_id, author_id)
			VALUES ($1, $2)`,
			bookID, authorID,
		)
		if err != nil {
			log.Printf("Error linking book and author: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	if err = tx.Commit(c); err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully", "book_id": bookID})
}
