package models

type Author struct {
	AuthorID  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
