package database

import (
	"os"

	"github.com/dee-d-dev/api/v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {

	dsn, exists := os.LookupEnv("DB_USER")

	if !exists {
		log.Fatal("DB_USER not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, models.Product{}, models.Image{}, models.Category{})
	return db
}

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{
		DB: db,
	}
}
