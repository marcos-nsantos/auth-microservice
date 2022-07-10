package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"log"
	"net/http"
	"strconv"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	IDUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	if err := userRepository.DeleteUser(uint(IDUint)); err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	helper.JSONResponse(w, http.StatusOK, "user deleted successfully")
}
