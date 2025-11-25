package repository

import "fmt"

type ProductRepositoryMysql struct {
}

func (p ProductRepositoryMysql) CreateProduct() {
	fmt.Println("ProductRepositoryMysql CreateProduct")
}

func (p ProductRepositoryMysql) GetProductByID() {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryMysql) GetProducts() {
	//TODO implement me
	panic("implement me")
}

func NewProductRepoMysql() ProductRepository {
	return ProductRepositoryMysql{}
}
