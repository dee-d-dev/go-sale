package main

import (
	"log"
	"net/http"

	"github.com/dee-d-dev/api/v1/router"
	"github.com/dee-d-dev/database"
)


func main() {

	db := database.Connect()	
	
	database.New(db)

	router := router.SetupRoutes()
	port := ":2001"

	log.Println("server started on localhost", port)

	http.ListenAndServe(port, router)

	
}

