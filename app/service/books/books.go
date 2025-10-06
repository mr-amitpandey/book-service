package ServiceBooks

import (
	"context"
	"fmt"

	repository "github.com/book-service/api/app/db/repository/books"
	dtosbooks "github.com/book-service/api/app/dtos/books"
	models "github.com/book-service/api/app/models/books"
	"github.com/google/uuid"
)

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookServiceInterface {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) CreateBook(ctx context.Context, book *dtosbooks.CreateBookRequest) (*dtosbooks.BookResponse, error) {
	bookID := uuid.New()
	bookModel := &models.Book{
		ID:    bookID,
		Name:  book.Name,
		Price: book.Price,
	}
	bookID, err := s.bookRepo.CreateBook(ctx, bookModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create book with details: %w", err)
	}
	return &dtosbooks.BookResponse{
		ID:    bookID.String(),
		Name:  book.Name,
		Price: book.Price,
	}, nil
}

func (s *bookService) UpdateBook(ctx context.Context, book *dtosbooks.UpdateBookRequest) (*dtosbooks.BookResponse, error) {
	// TODO: implement update book logic
	return nil, nil
}

func (s *bookService) DeleteBook(ctx context.Context, book_id string) error {
	// TODO: implement delete book logic
	return nil
}

func (s *bookService) GetBookByID(ctx context.Context, book_id string) (*dtosbooks.BookResponse, error) {
	bookUUID, err := uuid.Parse(book_id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %w", err)
	}
	book, err := s.bookRepo.GetBookByID(ctx, bookUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}
	return &dtosbooks.BookResponse{
		ID:    book.ID.String(),
		Name:  book.Name,
		Price: book.Price,
	}, nil
}
