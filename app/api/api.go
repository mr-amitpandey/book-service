package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routers struct {
	Router *gin.Engine
}

func (r *Routers) Init() {

	r.Router.GET("/v1/book-service/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Book Service is up and running",
		})
	})

	r.Router.StaticFile("/v1/book-service/swagger.yaml", "./docs/swagger.yaml")

	r.Router.GET("/v1/book-service/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/v1/book-service/swagger.yaml"),
	))

	v1 := r.Router.Group("/v1/book-service")

	// Book management routes
	bookGroup := BookGroup{
		RouterGroup: v1.Group("/books"),
	}

	// Employee management routes
	employeeGroup := EmployeeGroup{
		RouterGroup: v1.Group("/employees"),
	}

	// Initialize routes
	bookGroup.Init()
	employeeGroup.Init()

	defer func() {
		fmt.Println("Router has been initialized..")
	}()

}
