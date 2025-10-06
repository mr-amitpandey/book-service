package ServiceBooks

import (
	"context"

	dtosbooks "github.com/book-service/api/app/dtos/books"
)

type BookServiceInterface interface {
	CreateBook(ctx context.Context, book *dtosbooks.CreateBookRequest) (*dtosbooks.BookResponse, error)
	UpdateBook(ctx context.Context, book *dtosbooks.UpdateBookRequest) (*dtosbooks.BookResponse, error)
	DeleteBook(ctx context.Context, book_id string) error
	GetBookByID(ctx context.Context, book_id string) (*dtosbooks.BookResponse, error)
}
