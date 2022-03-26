package models

type Pet struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
