package dtosbooks

type CreateBookRequest struct {
	Name  string  `json:"name" binding:"required,min=2,max=100"`
	Price float32 `json:"price" binding:"required"` // Add more validations if needed
}

type UpdateBookRequest struct {
	Name  string  `json:"name" binding:"required,min=2,max=100"`
	Price float32 `json:"price" binding:"required"` // Add more validations if needed
}
