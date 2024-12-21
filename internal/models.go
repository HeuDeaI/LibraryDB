package internal

import "time"

type Book struct {
	BookID          int    `json:"book_id"`
	Title           string `json:"title"`
	PublicationYear int    `json:"publication_year"`
	Genre           string `json:"genre"`
}

type Author struct {
	AuthorID  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BookAuthor struct {
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}

type BookWithAuthors struct {
	BookID          int      `json:"book_id"`
	Title           string   `json:"title"`
	PublicationYear int      `json:"publication_year"`
	Genre           string   `json:"genre"`
	Authors         []string `json:"authors"`
}

type Reader struct {
	ReaderID    int    `json:"reader_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type Loan struct {
	LoanID     int        `json:"loan_id"`
	BookID     int        `json:"book_id"`
	ReaderID   int        `json:"reader_id"`
	IssueDate  time.Time  `json:"issue_date"`
	ReturnDate *time.Time `json:"return_date"`
}

type LoanRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Email     string `json:"email"`
	BookID    string `json:"book_id"`
}
