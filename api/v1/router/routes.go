package router

import (
	"github.com/gorilla/mux"
	"github.com/dee-d-dev/api/v1/controllers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// handler := &database.Handler{}

	v1.HandleFunc("/register", controllers.Register).Methods("POST")

	v1.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	v1.HandleFunc("/login", controllers.Login).Methods("POST")

	return r

	
}