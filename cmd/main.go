package main

import (
	"log"
	"orm-tests/internal/database"
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
	"orm-tests/internal/services"
	"orm-tests/internal/tests"
)

func main() {
	// Connect to PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Create tables, migrate
	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Create repositories
	categoryRepository := repositories.CreateCategoryRepository(db)
	productRepository := repositories.CreateProductRepository(db)

	// Create services
	categoryService := services.CreateCategoryService(categoryRepository)
	productService := services.CreateProductService(productRepository)

	// -1- Create and get test product and category
	tests.CreateAndGetProductAndCategory(categoryService, productService)

	// -2- Update test product
	tests.UpdateProduct(categoryService, productService) // Change IDs in this func

	// -3- Delete test category and products
	tests.DeleteCategoryAndProducts(categoryService, productService) // Change ID in this func
}
