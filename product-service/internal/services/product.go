package services

import (
	"product-service/internal/dto"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(productRequest dto.ProductRequest) (int, error) {

	// We will convert the dto to model
	product := models.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
	}
	id, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}
