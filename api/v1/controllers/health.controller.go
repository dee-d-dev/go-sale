package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status int    `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Message: "Server is up and running",
		Status: http.StatusOK,
	})
}