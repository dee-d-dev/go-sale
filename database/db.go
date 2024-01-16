package database

import (
	"os"

	"github.com/dee-d-dev/api/v1/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	db.AutoMigrate(&models.Users{})
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
