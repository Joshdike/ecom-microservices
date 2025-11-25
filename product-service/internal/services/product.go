package services

import "product-service/internal/repository"

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct() {
	s.productRepository.CreateProduct()
}
