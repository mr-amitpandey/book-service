package models

import "github.com/google/uuid"

type Employee struct {
	ID     uuid.UUID `json:"id"` // UUID for the book ID
	Name   string    `json:"name"`
	Salary float32   `json:"price"`
}
