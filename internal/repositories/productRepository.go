package repositories

import (
	"orm-tests/internal/models"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(product *models.Product) (uint, error)
	GetByID(id uint) (*models.Product, error)
	UpdateCategory(product *models.Product, newCategoryID uint) (uint, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func CreateProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) (uint, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) UpdateCategory(product *models.Product, newCategoryID uint) (uint, error) {
	result := r.db.Model(product).Update("CategoryID", newCategoryID)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}
