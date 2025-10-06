package api

import (
	booksHandler "github.com/book-service/api/app/handler/books"
	"github.com/gin-gonic/gin"
)

type BookGroup struct {
	RouterGroup *gin.RouterGroup
}

func (s *BookGroup) Init() {
	bookHandler := booksHandler.NewBookHandler()
	public := s.RouterGroup.Group("")

	public.POST("/create", bookHandler.CreateBook)
	public.PUT("/update/:id", bookHandler.UpdateBook)
	public.DELETE("/delete/:id", bookHandler.DeleteBook)
	public.GET("/getbyid/:id", bookHandler.GetBookByID)

}
