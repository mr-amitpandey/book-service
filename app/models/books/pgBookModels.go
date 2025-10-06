package models

import "github.com/google/uuid"

type Book struct {
	ID    uuid.UUID `json:"id"` // UUID for the book ID
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}
