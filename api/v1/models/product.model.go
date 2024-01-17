package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
	Images      []Image `json:"images" gorm:"foreignKey:ProductID"`
}

type Image struct {
	gorm.Model
	ProductID uint   `json:"-"`
	URL       string `json:"url"`
}
