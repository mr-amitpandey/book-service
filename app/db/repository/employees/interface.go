package repository

import (
	"context"

	models "github.com/book-service/api/app/models/employees"
	"github.com/google/uuid"
)

type EmployeeRepository interface {
	// Define methods for employee repository
	CreateEmployee(ctx context.Context, employee *models.Employee) (uuid.UUID, error)
	UpdateEmployee(ctx context.Context, employee *models.Employee) error
	DeleteEmployee(ctx context.Context, employee_id uuid.UUID) error
	GetEmployeeByID(ctx context.Context, employee_id uuid.UUID) (*models.Employee, error)
}
