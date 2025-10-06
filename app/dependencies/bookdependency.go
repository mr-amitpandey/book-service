package dependencies

import (
	"log"

	repository "github.com/book-service/api/app/db/repository/books"
	ServiceBooks "github.com/book-service/api/app/service/books"
)

type BookContainer struct {
	bookService    ServiceBooks.BookServiceInterface
	bookRepository repository.BookRepository
}

func (c *Container) GetBookRepository() repository.BookRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.BookContainer == nil {
		c.BookContainer = &BookContainer{}
	}

	if c.BookContainer.bookRepository == nil {
		c.BookContainer.bookRepository = repository.NewBookRepository(c.db)
	}
	return c.BookContainer.bookRepository
}

func (c *Container) GetBookService() ServiceBooks.BookServiceInterface {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.BookContainer == nil {
		c.BookContainer = &BookContainer{}
	}

	if c.BookContainer.bookService == nil {
		log.Println("🔧 Getting book repository...")
		// Don't call GetBookRepository() here - it would cause a deadlock
		// Instead, create the repository directly
		if c.BookContainer.bookRepository == nil {
			c.BookContainer.bookRepository = repository.NewBookRepository(c.db)
		}
		bookRepo := c.BookContainer.bookRepository
		log.Println("✅ Book repository obtained")

		log.Println("🔧 Creating book service...")
		c.BookContainer.bookService = ServiceBooks.NewBookService(bookRepo)
		log.Println("✅ Book service created")
	}
	return c.BookContainer.bookService
}
