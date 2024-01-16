package database

import (
	"os"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")

	}
	
	dsn := os.Getenv("DB_USER")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Users{})
	return db
}


type Handler struct{
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{
		DB: db,
	}
}

