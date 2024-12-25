package main

import (
	"fmt"
	"log"
	"orm-tests/internal/database"
	"orm-tests/internal/models"
	"orm-tests/internal/repositories"
	"orm-tests/internal/services"
)

func main() {
	// connect to PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Create tables
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

	// TESTS
	// 1. Create test category
	category, err := categoryService.CreateCategory("товары для дома")
	if err != nil {
		log.Fatal("failed to create category:", err)
	}
	fmt.Println(category)

	// 2. Create test product
	product, err := productService.CreateProduct("Веник new", 231.50, 1)
	if err != nil {
		log.Fatal("failed to create product:", err)
	}
	fmt.Println(product)

	// 3. Get test product
	product, err = productService.GetProductByID(product.ID)
	if err != nil {
		log.Fatal("failed to get product:", err)
	}
	fmt.Println(product)

	// 4. Get test category
	category, err = categoryService.GetCategoryByID(category.ID)
	if err != nil {
		log.Fatal("failed to get category:", err)
	}
	fmt.Println(category)

	// // Read
	// var userFromDb User
	// db.Preload("Posts").First(&userFromDb, "name = ?", "Alice")
	// println(userFromDb.Name)

	// // Update
	// db.Model(&userFromDb).Update("Name", "Alice Updated")

	// // Delete
	// db.Delete(&userFromDb.Posts)
	// db.Delete(&userFromDb)
}
