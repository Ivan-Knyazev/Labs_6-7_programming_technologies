package models

import "gorm.io/gorm"

// Model for Category
type Category struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

// Model for Product
type Product struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Price      float64
	CategoryID uint
}
