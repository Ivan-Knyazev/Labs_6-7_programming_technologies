package services

import (
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
)

type CategoryService interface {
	CreateCategory(name string) (*models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func CreateCategoryService(repository repositories.CategoryRepository) CategoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) CreateCategory(name string) (*models.Category, error) {
	category := &models.Category{Name: name}
	if _, err := s.repository.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repository.GetByID(id)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repository.Delete(id)
}
