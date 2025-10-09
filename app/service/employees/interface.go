package ServiceEmployees

import (
	"context"

	dtosemployees "github.com/book-service/api/app/dtos/employees"
)

type EmployeeServiceInterface interface {
	CreateEmployee(ctx context.Context, employee *dtosemployees.CreateEmployeeRequest) (*dtosemployees.EmployeeResponse, error)
	UpdateBook(ctx context.Context, employee *dtosemployees.UpdateEmployeeRequest) (*dtosemployees.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, employee_id string) error
	GetEmployeeByID(ctx context.Context, employee_id string) (*dtosemployees.EmployeeResponse, error)
}
