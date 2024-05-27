package personcontroller

import (
	"github.com/gorilla/mux"
	"go-jwt-eg/data/request"
	"go-jwt-eg/helpers"
	"go-jwt-eg/services"
	"net/http"
	"strconv"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	persons, err := services.FindAll(ctx)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Success", persons)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	persons, err := services.FindById(ctx, uint(id))
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Success", persons)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var reqPerson request.CreatePersonAdd
	if err := helpers.DecodeBody(w, r, &reqPerson); err != nil {
		return
	}
	defer r.Body.Close()

	if err := services.Create(reqPerson); err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusCreated, "Register successfully", nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var reqPerson request.UpdatePersonAdd
	if err := helpers.DecodeBody(w, r, &reqPerson); err != nil {
		return
	}
	defer r.Body.Close()

	reqPerson.ID = uint(id)
	resp := services.Update(reqPerson)

	helpers.Response(w, http.StatusOK, resp, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	resp := services.Delete(uint(id))

	helpers.Response(w, http.StatusOK, resp, nil)
}
