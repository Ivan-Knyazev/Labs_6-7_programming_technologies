package repositories

import (
	"orm-tests/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) (uint, error)
	GetByID(id uint) (*models.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func CreateCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// category := models.Category{Name: name, Products: []models.Product{}}
func (r *categoryRepository) Create(category *models.Category) (uint, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil
}

func (r *categoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	// Unscoped() - force delete
	result := r.db.Where("category_id = ?", id).Delete(&models.Product{})
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
