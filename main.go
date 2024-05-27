package main

import (
	"github.com/gorilla/mux"
	"go-jwt-eg/configs"
	"go-jwt-eg/routes"
	"log"
	"net/http"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.PersonRoutes(router)

	log.Println("Server running on port 8081")
	http.ListenAndServe(":8081", router)
}
