package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/model"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type ChangeUserPassword struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	IDUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var changeUserPassword ChangeUserPassword
	if err := json.NewDecoder(r.Body).Decode(&changeUserPassword); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	if changeUserPassword.NewPassword != changeUserPassword.ConfirmPassword {
		helper.JSONResponseWithError(w, http.StatusBadRequest, errors.New("passwords do not match"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	if _, err := userRepository.FindByID(uint(IDUint)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.JSONResponseWithError(w, http.StatusNotFound, errors.New("user not found"))
			return
		}
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	var user model.User
	user.ID = uint(IDUint)
	user.Password = changeUserPassword.NewPassword

	if err := user.HashPassword(); err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	if err := userRepository.UpdatePassword(&user); err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	helper.JSONResponse(w, http.StatusOK, "password updated successfully")
}
