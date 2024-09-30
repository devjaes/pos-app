package application

import (
	"github.com/devjaes/pos-app/internal/product/domain"
	"github.com/devjaes/pos-app/internal/product/infrastructure"
)

type ProductService struct {
	repo infrastructure.ProductRepository
}

func NewProductService(repo infrastructure.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name, description string, price float64, categoryID uint, supplierID uint) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		SupplierID:  supplierID,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}
