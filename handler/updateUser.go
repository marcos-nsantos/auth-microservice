package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/model"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"github.com/marcos-nsantos/e-commerce/auth-service/security"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type UserToUpdate struct {
	Name  string `json:"name" validate:"required,notblank,min=3,max=255"`
	Email string `json:"email" validate:"required,notblank,email"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	IDUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	userAuthID, err := security.ExtractUserID(r)
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusUnauthorized, err)
		return
	}

	if userAuthID != uint(IDUint) {
		helper.JSONResponseWithError(w, http.StatusForbidden, errors.New("you are not allowed to update this user"))
		return
	}

	var userToUpdate UserToUpdate
	if err := json.NewDecoder(r.Body).Decode(&userToUpdate); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.CheckForValidationErrMessages(userToUpdate); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	var user model.User
	user.ID = uint(IDUint)
	user.Name = userToUpdate.Name
	user.Email = userToUpdate.Email

	userRepository := repository.NewUserRepository(db)
	if err := userRepository.Update(&user); err != nil {
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
