package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/book-service/api/app/models/employees"
	"github.com/google/uuid"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee *models.Employee) (uuid.UUID, error) {

	query := `
		select sp_employee_create(
			$1, $2, $3
		)
	`
	var employeeID uuid.UUID
	err := r.db.QueryRowContext(ctx, query,
		employee.ID,
		employee.Name,
		employee.Salary,
	).Scan(&employeeID)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create employee with details: %w", err)
	}

	return employeeID, nil
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee *models.Employee) error {
	// TODO: implement update logic
	return nil
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employee_id uuid.UUID) error {
	// TODO: implement delete logic
	return nil
}

func (r *employeeRepository) GetEmployeeByID(ctx context.Context, employee_id uuid.UUID) (*models.Employee, error) {
	query := `SELECT 
		id, name, salary
	FROM fn_employee_get_by_id($1)`

	rows, err := r.db.QueryContext(ctx, query, employee_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("employee not found")
	}

	var employee models.Employee

	err = rows.Scan(
		&employee.ID,
		&employee.Name,
		&employee.Salary,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan employee data: %w", err)
	}

	return &employee, nil
}
