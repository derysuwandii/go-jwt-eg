package routes

import (
	"github.com/gorilla/mux"
	"go-jwt-eg/controllers/authcontroller"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/login", authcontroller.Login).Methods("POST")

}
