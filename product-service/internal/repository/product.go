package repository

type ProductRepository interface {
	CreateProduct()
	GetProductByID()
	GetProducts()
}
