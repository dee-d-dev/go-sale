package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/dee-d-dev/api/v1/router"
	"github.com/dee-d-dev/database"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	
	db := database.Connect()	
	
	database.New(db)

	router := router.SetupRoutes()
	port := ":2001"

	log.Println("server started on localhost", port)

	http.ListenAndServe(port, router)

	
}

