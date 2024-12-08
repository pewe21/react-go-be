package dto

type CreateProductDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

type UpdateProductDTO struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

type ProductDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}
