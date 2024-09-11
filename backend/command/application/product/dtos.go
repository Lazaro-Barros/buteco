package product

// CreateProductDTO é o DTO usado para criar um novo produto
type CreateProductDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Value       int    `json:"value" binding:"required,gt=0"`
	Visible     bool   `json:"visible"`
}

// UpdateProductDTO é o DTO usado para atualizar um produto existente
type UpdateProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value" binding:"omitempty,gt=0"`
	Visible     bool   `json:"visible"`
}
