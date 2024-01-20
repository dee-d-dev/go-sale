package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, "Error decoding json", http.StatusInternalServerError)
	}

	result := db.Create(&category)

	if result.Error != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(category)
}
