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

	v1.HandleFunc("/refresh", controllers.Refresh).Methods("POST")

	protectedWithMiddlware := http.HandlerFunc(controllers.ProtectedEndpoint)
	v1.Handle("/protected", middlewares.TokenMiddleWare(protectedWithMiddlware)).Methods("GET")

	createProduct := http.HandlerFunc(controllers.CreateProduct)
	v1.Handle("/products", middlewares.TokenMiddleWare(createProduct)).Methods("POST")

	getAllProducts := http.HandlerFunc(controllers.GetProducts)
	v1.Handle("/products", middlewares.TokenMiddleWare(getAllProducts)).Methods("GET")

	getAllUsers := http.HandlerFunc(controllers.GetUsers)
	v1.Handle("/users", middlewares.TokenMiddleWare(getAllUsers)).Methods("GET")

	getSingleUser := http.HandlerFunc(controllers.GetUser)
	v1.Handle("/users/{userId}", middlewares.TokenMiddleWare(getSingleUser)).Methods("GET")

	createCategory := http.HandlerFunc(controllers.CreateCategory)
	v1.Handle("/categories", middlewares.TokenMiddleWare(createCategory)).Methods("POST")

	return r

}
