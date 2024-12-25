package repositories

import (
	"orm-tests/internal/models"

	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	Create(category *models.Category) (uint, error)
	GetByID(id uint) (*models.Category, error)
	Delete(id uint) error
}

type CategoryRepository struct {
	db *gorm.DB
}

func CreateCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// category := models.Category{Name: name, Products: []models.Product{}}
func (r *CategoryRepository) Create(category *models.Category) (uint, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil
}

func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *CategoryRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
