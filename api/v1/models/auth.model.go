package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID        uint   `gorm: "primary_key" json:"id"` // uint is an unsigned integer
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Address   string `json:"address"`
	RToken    string `json:"rToken"`
}
