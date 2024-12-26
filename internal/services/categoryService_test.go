package services_test

import (
	"errors"
	"orm-tests/internal/models"
	"orm-tests/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCategoryRepository struct {
	mock.Mock
}

func (m *mockCategoryRepository) Create(category *models.Category) (uint, error) {
	category.ID = 1
	args := m.Called(category)
	return args.Get(0).(uint), args.Error(1)
}

func (m *mockCategoryRepository) GetByID(id uint) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockCategoryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCategorySuccess(t *testing.T) {
	mockRepo := &mockCategoryRepository{} // new(mockCategoryRepository)
	service := services.CreateCategoryService(mockRepo)

	mockRepo.On("Create", mock.AnythingOfType("*models.Category")).Return(uint(1), nil)

	createdCategory, err := service.CreateCategory("Test Category")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), createdCategory.ID)
	assert.Equal(t, "Test Category", createdCategory.Name)
	// mockRepo.AssertCalled(t, "Create", category)
}

func TestGetCategoryByIDSuccess(t *testing.T) {
	mockRepo := new(mockCategoryRepository)
	service := services.CreateCategoryService(mockRepo)
	var id uint = 1
	category := &models.Category{Name: "Test Category"}
	category.ID = id

	mockRepo.On("GetByID", id).Return(category, nil)

	foundCategory, err := service.GetCategoryByID(id)
	assert.NoError(t, err)
	assert.Equal(t, id, foundCategory.ID)
	assert.Equal(t, "Test Category", foundCategory.Name)
	mockRepo.AssertCalled(t, "GetByID", id)
}

func TestGetCategoryByIDNotFound(t *testing.T) {
	mockRepo := new(mockCategoryRepository)
	service := services.CreateCategoryService(mockRepo)
	var id uint = 2

	mockRepo.On("GetByID", id).Return(nil, errors.New("category not found"))

	_, err := service.GetCategoryByID(id)
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")
	mockRepo.AssertCalled(t, "GetByID", id)
}

func TestDeleteCategorySuccess(t *testing.T) {
	mockRepo := new(mockCategoryRepository)
	service := services.CreateCategoryService(mockRepo)
	var id uint = 1

	mockRepo.On("Delete", id).Return(nil)

	err := service.DeleteCategory(id)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", id)
}
