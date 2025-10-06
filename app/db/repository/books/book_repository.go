package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/book-service/api/app/models/books"
	"github.com/google/uuid"
)

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) CreateBook(ctx context.Context, book *models.Book) (uuid.UUID, error) {

	query := `
		select sp_books_create(
			$1, $2, $3
		)
	`
	var bookID uuid.UUID
	err := r.db.QueryRowContext(ctx, query,
		book.ID,
		book.Name,
		book.Price,
	).Scan(&bookID)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create book with details: %w", err)
	}

	return bookID, nil
}

func (r *bookRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	// TODO: implement update logic
	return nil
}

func (r *bookRepository) DeleteBook(ctx context.Context, book_id uuid.UUID) error {
	// TODO: implement delete logic
	return nil
}

func (r *bookRepository) GetBookByID(ctx context.Context, book_id uuid.UUID) (*models.Book, error) {
	query := `SELECT 
		id, name, price
	FROM fn_books_get_by_id($1)`

	rows, err := r.db.QueryContext(ctx, query, book_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("book not found")
	}

	var book models.Book

	err = rows.Scan(
		&book.ID,
		&book.Name,
		&book.Price,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan asset data: %w", err)
	}

	return &book, nil
}
