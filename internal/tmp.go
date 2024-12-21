package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SignUpReader(c *gin.Context, dbPool *pgxpool.Pool) {
	var reader Reader
	if err := c.BindJSON(&reader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := `
		INSERT INTO reader (first_name, last_name, phone_number, email)
		VALUES ($1, $2, $3, $4) RETURNING reader_id;
	`
	var readerID int
	err := dbPool.QueryRow(c, query, reader.FirstName, reader.LastName, reader.PhoneNumber, reader.Email).Scan(&readerID)
	if err != nil {
		log.Printf("Error inserting reader: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not sign up reader"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reader_id": readerID})
}

func LoanBook(c *gin.Context, dbPool *pgxpool.Pool) {
	var loan Loan
	if err := c.BindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := `
		INSERT INTO loan (book_id, reader_id, issue_date)
		VALUES ($1, $2, NOW()) RETURNING loan_id;
	`
	var loanID int
	err := dbPool.QueryRow(c, query, loan.BookID, loan.ReaderID).Scan(&loanID)
	if err != nil {
		log.Printf("Error creating loan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not loan book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loan_id": loanID})
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
