package models

type Book struct {
	BookID          int    `json:"book_id"`
	Title           string `json:"title"`
	PublicationYear int    `json:"publication_year"`
	Genre           string `json:"genre"`
}
