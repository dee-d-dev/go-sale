package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dee-d-dev/utils"
	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"

)


func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := database.Connect()

	var user models.Users


	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hashedPassword


	result := db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)

}

type LoginDetails struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string
	RefreshToken string
}

func Login(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	db := database.Connect()

	var credentials LoginDetails

	json.NewDecoder(r.Body).Decode(&credentials)

	var existingUser models.Users

	result := db.Where("email = ?", credentials.Email).First(&existingUser)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// password := r.FormValue("password")

	if utils.CheckPasswordHash(existingUser.Password, credentials.Password) {
		log.Fatal("Password does not match")
	}

	tokens, err := utils.GenerateToken(existingUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(LoginResponse{
		AccessToken: tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func ProtectedEndpoint (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Message: "You are authorized",
		Status: http.StatusOK,
	})
}