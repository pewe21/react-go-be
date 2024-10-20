package model

type ProductModel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}
