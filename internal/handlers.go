package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookWithAuthors struct {
	BookID          int      `json:"book_id"`
	Title           string   `json:"title"`
	PublicationYear int      `json:"publication_year"`
	Genre           string   `json:"genre"`
	Authors         []string `json:"authors"`
}

func GetBooksWithAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	query := `
		SELECT 
			b.book_id, b.title, b.publication_year, b.genre, 
			STRING_AGG(a.first_name || ' ' || a.last_name, ', ') AS authors
		FROM 
			book AS b
		LEFT JOIN 
			bookauthor AS ba ON b.book_id = ba.book_id
		LEFT JOIN 
			author AS a ON ba.author_id = a.author_id
		GROUP BY 
			b.book_id
		ORDER BY 
			b.title;
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

func GetBooks(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT book_id, title, publication_year, genre FROM book")
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.PublicationYear, &book.Genre); err != nil {
			log.Printf("Error scanning book: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT author_id, first_name, last_name FROM author")
	if err != nil {
		log.Printf("Error fetching authors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.AuthorID, &author.FirstName, &author.LastName); err != nil {
			log.Printf("Error scanning author: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func GetReaders(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT reader_id, first_name, last_name, phone_number, email FROM reader")
	if err != nil {
		log.Printf("Error fetching readers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var readers []Reader
	for rows.Next() {
		var reader Reader
		if err := rows.Scan(&reader.ReaderID, &reader.FirstName, &reader.LastName, &reader.PhoneNumber, &reader.Email); err != nil {
			log.Printf("Error scanning reader: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		readers = append(readers, reader)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, readers)
}

func GetBookAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT book_id, author_id FROM bookauthor")
	if err != nil {
		log.Printf("Error fetching book-author relationships: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var bookAuthors []BookAuthor
	for rows.Next() {
		var bookAuthor BookAuthor
		if err := rows.Scan(&bookAuthor.BookID, &bookAuthor.AuthorID); err != nil {
			log.Printf("Error scanning book-author relationship: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		bookAuthors = append(bookAuthors, bookAuthor)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, bookAuthors)
}

func GetLoans(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT loan_id, book_id, reader_id, issue_date, return_date FROM loan")
	if err != nil {
		log.Printf("Error fetching loans: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var loans []Loan
	for rows.Next() {
		var loan Loan
		if err := rows.Scan(&loan.LoanID, &loan.BookID, &loan.ReaderID, &loan.IssueDate, &loan.ReturnDate); err != nil {
			log.Printf("Error scanning loan: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, loans)
}
