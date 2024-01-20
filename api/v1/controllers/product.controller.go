package controllers

import (
	"encoding/json"
	"log"

	// "log"
	"net/http"

	"github.com/dee-d-dev/api/v1/models"
	"github.com/dee-d-dev/database"
	"github.com/dee-d-dev/utils"
	"gorm.io/gorm/clause"
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
		Name:        productDetails.Name,
		Description: productDetails.Description,
		Category:    productDetails.Category,
		Price:       productDetails.Price,
		Stock:       productDetails.Stock,
		Brand:       productDetails.Brand,
		Color:       productDetails.Color,
		Size:        productDetails.Size,
		Images:      productDetails.Images,
		MerchantID:  user.ID,
		Merchant:    user,
	}
	db.Create(&product)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
	// json.NewEncoder(w).Encode(loggedInUserEmail)

}

// type ProductResponse struct {
// 	ID          uint           `json:"id"`
// 	Name        string         `json:"name"`
// 	Description string         `json:"description"`
// 	Price       float64        `json:"price"`
// 	Stock       int            `json:"stock"`
// 	Category    string         `json:"category"`
// 	Brand       string         `json:"brand"`
// 	Color       *string        `json:"color"`
// 	Size        string         `json:"size"`
// 	Images      []models.Image `json:"images"`
// }

type ProductResponse struct {
	products []models.Product
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	var products []models.Product

	if err := db.Preload(clause.Associations).Find(&products).Error; err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		log.Println("Error fetching products:", err)
		return
	}

	// productsJSON, err := json.Marshal(products)

	// if err != nil {
	// 	http.Error(w, "Error encoding products to JSON", http.StatusInternalServerError)
	//     log.Println("Error encoding products to JSON:", err)
	//     return
	// }

	// var productResponses []ProductResponse

	// for _, product := range products {
	//     productResponses = append(productResponses, ProductResponse{
	//         ID:   product.ID,
	//         Name: product.Name,
	// 		Description: product.Description,
	// 		Price: product.Price,
	// 		Stock: product.Stock,
	// 		Category: product.Category,
	// 		Brand: product.Brand,
	// 		Color: product.Color,
	// 		Size: product.Size,
	// 		Images: product.Images,
	//     })
	// }
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	// w.Write(productsJSON)
	json.NewEncoder(w).Encode(products)

}

func CreateCategory(w http.ResponseWriter, r *http.Request){
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
