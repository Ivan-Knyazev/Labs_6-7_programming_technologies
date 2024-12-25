package services

import (
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
)

type ProductService interface {
	CreateProduct(name string, price float64, categoryID uint) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	GetProductByCategory(categoryID uint) (*[]*models.Product, error)
	UpdateCategoryForProduct(productID uint, newCategoryID uint) (*models.Product, error)
}

type productService struct {
	repository repositories.ProductRepository
}

func CreateProductService(repository repositories.ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (s *productService) CreateProduct(name string, price float64, categoryID uint) (*models.Product, error) {
	product := &models.Product{Name: name, Price: price, CategoryID: categoryID}
	if _, err := s.repository.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.repository.GetByID(id)
}

func (s *productService) GetProductByCategory(categoryID uint) (*[]*models.Product, error) {
	return s.repository.GetByCategory(categoryID)
}

func (s *productService) UpdateCategoryForProduct(productID uint, newCategoryID uint) (*models.Product, error) {
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
