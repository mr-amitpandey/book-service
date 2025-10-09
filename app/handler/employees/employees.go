package employeesHandler

import (
	"github.com/book-service/api/app/dependencies"
	dtosemployees "github.com/book-service/api/app/dtos/employees"
	"github.com/book-service/api/app/helper/response"
	ServiceEmployees "github.com/book-service/api/app/service/employees"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService ServiceEmployees.EmployeeServiceInterface
}

func NewEmployeeHandler() *EmployeeHandler {
	container := dependencies.GetContainer()
	employeeService := container.GetEmployeeService()
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	// TODO: Implement create book handler
	var req dtosemployees.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
		return
	}
	employeeResponse, err := h.employeeService.CreateEmployee(c, &req)
	if err != nil {
		response.SendBadRequestResponse(c, err.Error())
		return
	}
	response.SendSuccessResponse(c, employeeResponse)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	// TODO: Implement update book handler
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	// TODO: Implement delete book handler
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		response.SendBadRequestResponse(c, "Employee ID is required")
		return
	}

	employee, err := h.employeeService.GetEmployeeByID(c, employeeID)
	if err != nil {
		response.SendNotFoundResponse(c, err.Error())
		return
	}

	response.SendSuccessResponse(c, employee)
}
