package services_test

import (
	"errors"
	"orm-tests/internal/models"
	"orm-tests/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockProductRepository struct {
	mock.Mock
}

func (m *mockProductRepository) Create(product *models.Product) (uint, error) {
	product.ID = 1
	args := m.Called(product)
	return args.Get(0).(uint), args.Error(1)
}

func (m *mockProductRepository) GetByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockProductRepository) GetByCategory(categoryID uint) (*[]*models.Product, error) {
	args := m.Called(categoryID)
	if args.Get(0) != nil {
		return args.Get(0).(*[]*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockProductRepository) UpdateCategory(product *models.Product, newCategoryID uint) (uint, error) {
	args := m.Called(product, newCategoryID)
	return args.Get(0).(uint), args.Error(1)
}

func TestCreateProductSuccess(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)

	mockRepo.On("Create", mock.AnythingOfType("*models.Product")).Return(uint(1), nil)

	createdProduct, err := service.CreateProduct("brush", 234.56, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), createdProduct.ID)
	assert.Equal(t, "brush", createdProduct.Name)
	assert.Equal(t, 234.56, createdProduct.Price)
	assert.Equal(t, uint(1), createdProduct.CategoryID)
}

func TestGetProductByIDSuccess(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	var id uint = 1
	product := &models.Product{Name: "brush", Price: 234.56, CategoryID: 1}
	product.ID = id

	mockRepo.On("GetByID", id).Return(product, nil)

	foundProduct, err := service.GetProductByID(id)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), foundProduct.ID)
	assert.Equal(t, "brush", foundProduct.Name)
	assert.Equal(t, 234.56, foundProduct.Price)
	assert.Equal(t, uint(1), foundProduct.CategoryID)
	mockRepo.AssertCalled(t, "GetByID", id)
}

func TestGetProductByIDNotFound(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	var id uint = 2

	mockRepo.On("GetByID", id).Return(nil, errors.New("product not found"))

	_, err := service.GetProductByID(id)
	assert.Error(t, err)
	assert.EqualError(t, err, "product not found")
	mockRepo.AssertCalled(t, "GetByID", id)
}

func TestGetProductByCategorySuccess(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	products := make([]*models.Product, 0)
	for i := 0; i < 3; i++ {
		products = append(products, &models.Product{Name: "brush", Price: 234.56 + float64(i), CategoryID: 1})
		products[i].ID = uint(i)
	}
	var categoryID uint = 1

	mockRepo.On("GetByCategory", categoryID).Return(&products, nil)

	foundProducts, err := service.GetProductByCategory(categoryID)
	assert.NoError(t, err)
	for id, product := range *foundProducts {
		assert.Equal(t, uint(id), product.ID)
		assert.Equal(t, "brush", product.Name)
		assert.Equal(t, 234.56+float64(id), product.Price)
		assert.Equal(t, uint(1), product.CategoryID)
	}
	mockRepo.AssertCalled(t, "GetByCategory", categoryID)
}

func TestGetProductByCategoryNotFount(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	var categoryID uint = 5

	mockRepo.On("GetByCategory", categoryID).Return(nil, errors.New("products for this category not found"))

	_, err := service.GetProductByCategory(categoryID)
	assert.Error(t, err)
	assert.EqualError(t, err, "products for this category not found")
	mockRepo.AssertCalled(t, "GetByCategory", categoryID)
}

func TestUpdateCategoryForProductSuccess(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	var productID uint = 1
	var newCategoryID uint = 3
	product := &models.Product{Name: "brush", Price: 234.56, CategoryID: 1}
	product.ID = productID

	mockRepo.On("GetByID", productID).Return(product, nil)
	product.CategoryID = newCategoryID
	mockRepo.On("UpdateCategory", mock.AnythingOfType("*models.Product"), newCategoryID).Return(newCategoryID, nil)

	foundProduct, err := service.UpdateCategoryForProduct(productID, newCategoryID)
	assert.NoError(t, err)
	assert.Equal(t, productID, foundProduct.ID)
	assert.Equal(t, "brush", foundProduct.Name)
	assert.Equal(t, 234.56, foundProduct.Price)
	assert.Equal(t, uint(3), foundProduct.CategoryID)
	mockRepo.AssertCalled(t, "GetByID", productID)
	mockRepo.AssertCalled(t, "UpdateCategory", mock.AnythingOfType("*models.Product"), newCategoryID)
}

func TestUpdateCategoryForProductNotFound(t *testing.T) {
	mockRepo := new(mockProductRepository)
	service := services.CreateProductService(mockRepo)
	var productID uint = 1
	var newCategoryID uint = 3
	// product := &models.Product{Name: "brush", Price: 234.56, CategoryID: 1}
	// product.ID = productID

	mockRepo.On("GetByID", productID).Return(nil, errors.New("product not found"))
	mockRepo.On("UpdateCategory", mock.AnythingOfType("*models.Product"), newCategoryID).Return(newCategoryID, nil)

	_, err := service.UpdateCategoryForProduct(productID, newCategoryID)
	assert.Error(t, err)
	assert.EqualError(t, err, "product not found")
	mockRepo.AssertCalled(t, "GetByID", productID)
	mockRepo.AssertNotCalled(t, "UpdateCategory", mock.AnythingOfType("*models.Product"), newCategoryID)
}
