package services

import (
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
)

type CategoryServiceInterface interface {
	CreateCategory(name string) (*models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	DeleteCategory(id uint) error
}

type CategoryService struct {
	repository *repositories.CategoryRepository
}

func CreateCategoryService(repository *repositories.CategoryRepository) CategoryService {
	return CategoryService{repository: repository}
}

func (s *CategoryService) CreateCategory(name string) (*models.Category, error) {
	category := &models.Category{Name: name}
	if _, err := s.repository.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repository.GetByID(id)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repository.Delete(id)
}
