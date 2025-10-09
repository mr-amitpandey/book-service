package dtosemployees

type EmployeeResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name" binding:"required,min=2,max=100"`
	Salary float32 `json:"salary"`
}
