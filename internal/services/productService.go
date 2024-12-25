package services

import (
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
)

type ProductServiceInterface interface {
	CreateProduct(name string, price float64, categoryID uint) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateCategoryForProduct(productID uint, newCategoryID uint) (*models.Product, error)
}

type ProductService struct {
	repository *repositories.ProductRepository
}

func CreateProductService(repository *repositories.ProductRepository) ProductService {
	return ProductService{repository: repository}
}

func (s *ProductService) CreateProduct(name string, price float64, categoryID uint) (*models.Product, error) {
	product := &models.Product{Name: name, Price: price, CategoryID: categoryID}
	if _, err := s.repository.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.repository.GetByID(id)
}

func (s *ProductService) UpdateCategoryForProduct(productID uint, newCategoryID uint) (*models.Product, error) {
	var product *models.Product
	product, err := s.repository.GetByID(productID)
	if err != nil {
		return nil, err
	}
	if _, err := s.repository.UpdateCategory(product, newCategoryID); err != nil {
		return nil, err
	}
	return product, nil
}
