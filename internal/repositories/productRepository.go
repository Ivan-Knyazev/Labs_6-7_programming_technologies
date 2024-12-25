package repositories

import (
	"orm-tests/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) (uint, error)
	GetByID(id uint) (*models.Product, error)
	GetByCategory(categoryID uint) (*[]*models.Product, error)
	UpdateCategory(product *models.Product, newCategoryID uint) (uint, error)
}

type productRepository struct {
	db *gorm.DB
}

func CreateProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) (uint, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) GetByCategory(categoryID uint) (*[]*models.Product, error) {
	products := make([]*models.Product, 0)
	result := r.db.Where("category_id = ?", categoryID).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (r *productRepository) UpdateCategory(product *models.Product, newCategoryID uint) (uint, error) {
	result := r.db.Model(product).Update("CategoryID", newCategoryID)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}
