package dtosemployees

type CreateEmployeeRequest struct {
	Name   string  `json:"name" binding:"required,min=2,max=100"`
	Salary float32 `json:"salary" binding:"required"` // Add more validations if needed
}

type UpdateEmployeeRequest struct {
	Name   string  `json:"name" binding:"required,min=2,max=100"`
	Salary float32 `json:"salary" binding:"required"` // Add more validations if needed
}
