package booksHandler

import (
	"github.com/book-service/api/app/dependencies"
	dtosbooks "github.com/book-service/api/app/dtos/books"
	"github.com/book-service/api/app/helper/response"
	ServiceBooks "github.com/book-service/api/app/service/books"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService ServiceBooks.BookServiceInterface
}

func NewBookHandler() *BookHandler {
	container := dependencies.GetContainer()
	bookService := container.GetBookService()
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	// TODO: Implement create book handler
	var req dtosbooks.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
		return
	}
	bookResponse, err := h.bookService.CreateBook(c, &req)
	if err != nil {
		response.SendBadRequestResponse(c, err.Error())
		return
	}
	response.SendSuccessResponse(c, bookResponse)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	// TODO: Implement update book handler
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	// TODO: Implement delete book handler
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	if bookID == "" {
		response.SendBadRequestResponse(c, "Book ID is required")
		return
	}

	book, err := h.bookService.GetBookByID(c, bookID)
	if err != nil {
		response.SendNotFoundResponse(c, err.Error())
		return
	}

	response.SendSuccessResponse(c, book)
}
