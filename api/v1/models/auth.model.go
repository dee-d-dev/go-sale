package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm: "primary_key" json:"id"` // uint is an unsigned integer
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index;not null"`
	Password  string `json:"-" gorm:"not null"`
	Address   string `json:"address" gorm:"not null"`
	RToken    string `json:"rToken"`
	Products  []Product `json:"products" gorm:"foreignkey:MerchantID"`
}
