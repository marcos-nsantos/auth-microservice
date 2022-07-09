package handler

import (
	"encoding/json"
	"errors"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/model"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.CheckForValidationErrMessages(user); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	if err := user.HashPassword(); err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	if err := userRepository.Create(&user); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	user.Password = ""

	helper.JSONResponse(w, http.StatusCreated, user)
}
