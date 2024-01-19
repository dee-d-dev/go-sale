package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `gorm: "primary_key" json:"id"` 
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Color       *string  `json:"color"`
	Size        string  `json:"size"`
	Images      []Image `json:"images" gorm:"foreignKey:ProductID"`
	MerchantID  uint    `json:"merchant_id" gorm:"not null"` // Foreign key referencing User's ID
	Merchant    User    `json:"merchant" gorm:"foreignKey:MerchantID"`
}

type Image struct {
	gorm.Model
	ProductID uint   `json:"-"`
	URL       string `json:"url"`
}
