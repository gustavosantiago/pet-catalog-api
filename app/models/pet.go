package models

import "time"

type Pet struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Breed       string    `json:"breed"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
