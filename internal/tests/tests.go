package tests

import (
	"fmt"
	"log"
	"orm-tests/internal/services"
	"orm-tests/internal/utils"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

func CreateAndGetProductAndCategory(categoryService services.CategoryService, productService services.ProductService) {
	// Generating a random number for different values in the DB to run the app many times
	rand.Seed(uint64(time.Now().UnixNano()))
	randomInt := strconv.Itoa(rand.Intn(100))

	// TESTS
	// 1. Create test category
	fmt.Println("1. Create test category")
	category, err := categoryService.CreateCategory("Электроника " + randomInt)
	if err != nil {
		log.Fatal("failed to create category:", err)
	}
	utils.SerializeAndPrint(*category)

	// 2. Create test product
	fmt.Println("2. Create test product")
	product, err := productService.CreateProduct("Веник new "+randomInt, 231.50+float64(rand.Intn(100)), 2)
	if err != nil {
		log.Fatal("failed to create product:", err)
	}
	utils.SerializeAndPrint(*product)

	// 3. Get test product
	fmt.Println("3. Get test product")
	product, err = productService.GetProductByID(product.ID)
	if err != nil {
		log.Fatal("failed to get product:", err)
	}
	utils.SerializeAndPrint(*product)

	// 4. Get test category
	fmt.Println("4. Get test category")
	category, err = categoryService.GetCategoryByID(category.ID)
	if err != nil {
		log.Fatal("failed to get category:", err)
	}
	utils.SerializeAndPrint(*category)

	// 5. Get products by category
	fmt.Println("5. Get products by category")
	products, err := productService.GetProductByCategory(2)
	if err != nil {
		log.Fatal("failed to get products by category:", err)
	}
	fmt.Println("Products by category 1:")
	for _, value := range *products {
		utils.SerializeAndPrint(*value)
	}
}

func UpdateProduct(categoryService services.CategoryService, productService services.ProductService) {
	// 6. Update category for product
	fmt.Println("6. Update category for product")
	product, err := productService.UpdateCategoryForProduct(5, 5)
	if err != nil {
		log.Fatal("failed to update category for product:", err)
	}
	utils.SerializeAndPrint(*product)
}

func DeleteCategoryAndProducts(categoryService services.CategoryService, productService services.ProductService) {
	// 7. Delete products and their category
	fmt.Println("7. Delete products and their category")
	err := categoryService.DeleteCategory(5)
	if err != nil {
		log.Fatal("failed to D]delete products and their category:", err)
	}
}
