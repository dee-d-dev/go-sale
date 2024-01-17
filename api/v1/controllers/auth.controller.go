package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"
	"github.com/dee-d-dev/utils"

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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

func Login(w http.ResponseWriter, r *http.Request) {
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

	accessToken, err := utils.GenerateAccessToken(existingUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	refreshToken, err := utils.GenerateRefreshToken(existingUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	resultT := db.Model(&models.Users{}).Where("email = ?", existingUser.Email).Update("r_token", refreshToken)

	if resultT.Error != nil {
		log.Fatal(resultT.Error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Message: "You are authorized",
		Status:  http.StatusOK,
	})
}

type RefreshToken struct{
	Token string `json:"refresh_token"`
}

type RefreshTokenResponse struct{
	AccessToken string `json:"access_token"`
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	var rt RefreshToken
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&rt)


	if err != nil {
		log.Fatal(err)
		
	}

	err = db.Where("r_token = ?", rt.Token).First(&user).Error

	if err != nil {
		
		log.Fatal(err)
	}

	// newAccessTokenExpiration := time.Now().Add(15 * time.Minute)
	
	newAcessToken, err := utils.GenerateAccessToken(user.Email)

	if err != nil {
		log.Fatal(err)
	}



	// if(rt != r.Header.Get("Authorization")) {
	// 	log.Fatal("Invalid refresh token")
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&RefreshTokenResponse{
		AccessToken: newAcessToken,
	})

}
