package models

import "time"

type Loan struct {
	LoanID     int        `json:"loan_id"`
	BookID     int        `json:"book_id"`
	ReaderID   int        `json:"reader_id"`
	IssueDate  time.Time  `json:"issue_date"`
	ReturnDate *time.Time `json:"return_date"`
}
