package repository

import (
	"context"

	models "github.com/book-service/api/app/models/books"
	"github.com/google/uuid"
)

type BookRepository interface {
	// Define methods for book repository
	CreateBook(ctx context.Context, book *models.Book) (uuid.UUID, error)
	UpdateBook(ctx context.Context, book *models.Book) error
	DeleteBook(ctx context.Context, book_id uuid.UUID) error
	GetBookByID(ctx context.Context, book_id uuid.UUID) (*models.Book, error)
}
