package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var users []models.User
	db := database.Connect()

	if err := db.Preload(clause.Associations).Find(&users).Error; err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := database.Connect()

	var user models.User

	vars := mux.Vars(r)

	userId := vars["userId"]

	result := db.Preload(clause.Associations).Where("id = ?", userId).First(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)

}
