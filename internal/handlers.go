package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetBooks(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT book_id, title, publication_year, genre FROM book")
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.PublicationYear, &book.Genre); err != nil {
			log.Printf("Error scanning book: %v", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "books.html", gin.H{"Books": books})
}

func GetAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT author_id, first_name, last_name FROM author")
	if err != nil {
		log.Printf("Error fetching authors: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.AuthorID, &author.FirstName, &author.LastName); err != nil {
			log.Printf("Error scanning author: %v", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "authors.html", gin.H{"Authors": authors})
}

func GetReaders(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT reader_id, first_name, last_name, phone_number, email FROM reader")
	if err != nil {
		log.Printf("Error fetching readers: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var readers []Reader
	for rows.Next() {
		var reader Reader
		if err := rows.Scan(&reader.ReaderID, &reader.FirstName, &reader.LastName, &reader.PhoneNumber, &reader.Email); err != nil {
			log.Printf("Error scanning reader: %v", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		readers = append(readers, reader)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "readers.html", gin.H{"Readers": readers})
}

func GetBookAuthors(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT book_id, author_id FROM bookauthor")
	if err != nil {
		log.Printf("Error fetching book-author relationships: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var bookAuthors []BookAuthor
	for rows.Next() {
		var bookAuthor BookAuthor
		if err := rows.Scan(&bookAuthor.BookID, &bookAuthor.AuthorID); err != nil {
			log.Printf("Error scanning book-author relationship: %v", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		bookAuthors = append(bookAuthors, bookAuthor)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "bookauthors.html", gin.H{"BookAuthors": bookAuthors})
}

func GetLoans(c *gin.Context, dbPool *pgxpool.Pool) {
	rows, err := dbPool.Query(c, "SELECT loan_id, book_id, reader_id, issue_date, return_date FROM loan")
	if err != nil {
		log.Printf("Error fetching loans: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var loans []Loan
	for rows.Next() {
		var loan Loan
		if err := rows.Scan(&loan.LoanID, &loan.BookID, &loan.ReaderID, &loan.IssueDate, &loan.ReturnDate); err != nil {
			log.Printf("Error scanning loan: %v", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "loans.html", gin.H{"Loans": loans})
}
