package usercontroller

import (
	"go-jwt-eg/helpers"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	helpers.Response(w, http.StatusOK, "My Profile", user)
}
func GelAll(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	helpers.Response(w, http.StatusOK, "My Profile", user)
}
