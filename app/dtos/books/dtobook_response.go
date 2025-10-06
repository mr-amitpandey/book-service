package dtosbooks

type BookResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name" binding:"required,min=2,max=100"`
	Price float32 `json:"price"`
}
