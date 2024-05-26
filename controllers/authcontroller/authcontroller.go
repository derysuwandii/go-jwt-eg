package authcontroller

import (
	"go-jwt-eg/configs"
	"go-jwt-eg/helpers"
	"go-jwt-eg/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := helpers.DecodeBody(w, r, &register); err != nil {
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, http.StatusBadRequest, "Password not match", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusCreated, "Register successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := helpers.DecodeBody(w, r, &login); err != nil {
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email=?", login.Email).Error; err != nil {
		helpers.Response(w, http.StatusNotFound, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, http.StatusNotFound, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	loginRes := models.LoginRes{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	helpers.Response(w, http.StatusOK, "Successfuly Login...", loginRes)
}
