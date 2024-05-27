package routes

import (
	"github.com/gorilla/mux"
	"go-jwt-eg/controllers/personcontroller"
	"go-jwt-eg/middleware"
)

func PersonRoutes(r *mux.Router) {
	router := r.PathPrefix("/persons").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("", personcontroller.FindAll).Methods("GET")
	router.HandleFunc("/{id}", personcontroller.FindById).Methods("GET")
	router.HandleFunc("/add", personcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}", personcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}", personcontroller.Delete).Methods("DELETE")

}
