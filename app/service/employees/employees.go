package ServiceEmployees

import (
	"context"
	"fmt"

	repository "github.com/book-service/api/app/db/repository/employees"
	dtosemployees "github.com/book-service/api/app/dtos/employees"

	models "github.com/book-service/api/app/models/employees"
	"github.com/google/uuid"
)

type employeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeServiceInterface {
	return &employeeService{
		employeeRepo: employeeRepo,
	}
}

func (s *employeeService) CreateEmployee(ctx context.Context, employee *dtosemployees.CreateEmployeeRequest) (*dtosemployees.EmployeeResponse, error) {
	employeeID := uuid.New()
	employeeModel := &models.Employee{
		ID:     employeeID,
		Name:   employee.Name,
		Salary: employee.Salary,
	}
	employeeID, err := s.employeeRepo.CreateEmployee(ctx, employeeModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee with details: %w", err)
	}
	return &dtosemployees.EmployeeResponse{
		ID:     employeeID.String(),
		Name:   employee.Name,
		Salary: employee.Salary,
	}, nil
}

func (s *employeeService) UpdateBook(ctx context.Context, employee *dtosemployees.UpdateEmployeeRequest) (*dtosemployees.EmployeeResponse, error) {
	// TODO: implement update book logic
	return nil, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, employee_id string) error {
	// TODO: implement delete book logic
	return nil
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, employee_id string) (*dtosemployees.EmployeeResponse, error) {
	employeeUUID, err := uuid.Parse(employee_id)
	if err != nil {
		return nil, fmt.Errorf("invalid employee ID: %w", err)
	}
	employee, err := s.employeeRepo.GetEmployeeByID(ctx, employeeUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee: %w", err)
	}
	return &dtosemployees.EmployeeResponse{
		ID:     employee.ID.String(),
		Name:   employee.Name,
		Salary: employee.Salary,
	}, nil
}
