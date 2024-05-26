package routes

import (
	"github.com/gorilla/mux"
	"go-jwt-eg/controllers/usercontroller"
	"go-jwt-eg/middleware"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/me", usercontroller.Me).Methods("GET")

}
