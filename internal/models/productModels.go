package models

import "gorm.io/gorm"

// Model for Category
type Category struct {
	gorm.Model
	Name     string    `gorm:"not null" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

// Model for Product
type Product struct {
	gorm.Model
	Name       string  `gorm:"not null" json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"categoryID"`
}
