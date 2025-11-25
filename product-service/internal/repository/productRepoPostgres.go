package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryPostgres struct {
	db *sqlx.DB
}

func (p *ProductRepositoryPostgres) CreateProduct() {
	fmt.Println("ProductRepositoryPostgres CreateProduct")
	fmt.Println("ProductRepositoryPostgres CreateProduct")
}

func (p *ProductRepositoryPostgres) GetProductByID() {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepositoryPostgres) GetProducts() {
	//TODO implement me
	panic("implement me")
}

func NewProductRepositoryPostgres() ProductRepository {
	conn := sqlx.MustConnect("postgres", "postgres:5432?sslmode=disable")
	return &ProductRepositoryPostgres{
		conn,
	}
}
