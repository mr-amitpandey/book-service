package api

import (
	employeesHandler "github.com/book-service/api/app/handler/employees"
	"github.com/gin-gonic/gin"
)

type EmployeeGroup struct {
	RouterGroup *gin.RouterGroup
}

func (s *EmployeeGroup) Init() {
	employeeHandler := employeesHandler.NewEmployeeHandler()

	s.RouterGroup.POST("/create", employeeHandler.CreateEmployee)
	s.RouterGroup.PUT("/update/:id", employeeHandler.UpdateEmployee)
	s.RouterGroup.DELETE("/delete/:id", employeeHandler.DeleteEmployee)
	s.RouterGroup.GET("/get/:id", employeeHandler.GetEmployeeByID)

}
