package router

import (
	"net/http"

	"github.com/dee-d-dev/api/v1/controllers"
	"github.com/dee-d-dev/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// handler := &database.Handler{}

	v1.HandleFunc("/register", controllers.Register).Methods("POST")

	v1.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	v1.HandleFunc("/login", controllers.Login).Methods("POST")

	protectedWithMiddlware := http.HandlerFunc(controllers.ProtectedEndpoint)
	v1.Handle("/protected", middlewares.TokenMiddleWare(protectedWithMiddlware)).Methods("GET")

	// v1.Handle("/refresh", controllers.Refresh).Methods("POST")
	return r

	
}