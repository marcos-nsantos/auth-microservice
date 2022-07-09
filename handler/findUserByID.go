package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	IDUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.JSONResponse(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByID(uint(IDUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.JSONResponseWithError(w, http.StatusNotFound, errors.New("user not found"))
			return
		}
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	helper.JSONResponse(w, http.StatusOK, user)
}
