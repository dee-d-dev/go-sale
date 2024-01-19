package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"
	"github.com/dee-d-dev/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := database.Connect()

	var productDetails models.Product
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&productDetails)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	loggedInUserEmail, err := utils.GetLoggedInUser(r)

	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	if err := db.Where("email = ?", loggedInUserEmail).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	product := models.Product{
		Name: productDetails.Name,
		Description: productDetails.Description,
		Category: productDetails.Category,
		Price: productDetails.Price,
		Stock: productDetails.Stock,
		Brand: productDetails.Brand,
		Color: productDetails.Color,
		Size: productDetails.Size,
		Images: productDetails.Images,
		MerchantID: user.ID,
		Merchant: user,
	}
	db.Create(&product)


	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
	// json.NewEncoder(w).Encode(loggedInUserEmail)

}
