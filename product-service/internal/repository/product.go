package repository

import "product-service/internal/models"

type ProductRepository interface {
	CreateProduct(product models.Product) (int, error)
	GetProductByID()
	GetProducts()
}
