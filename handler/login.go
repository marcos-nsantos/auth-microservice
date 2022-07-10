package handler

import (
	"encoding/json"
	"errors"
	"github.com/marcos-nsantos/e-commerce/auth-service/database"
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/repository"
	"github.com/marcos-nsantos/e-commerce/auth-service/security"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type login struct {
	Email    string `json:"email" validate:"required,email,notblank"`
	Password string `json:"password" validate:"required,notblank"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.CheckForValidationErrMessages(login); err != nil {
		helper.JSONResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByEmail(login.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.JSONResponseWithError(w, http.StatusNotFound, errors.New("email or password is incorrect"))
			return
		}
		helper.JSONResponseWithError(w, http.StatusInternalServerError, errors.New("something went wrong"))
		log.Println(err)
		return
	}

	if err := user.CheckPassword(login.Password); err != nil {
		helper.JSONResponseWithError(w, http.StatusUnauthorized, errors.New("email or password is incorrect"))
		return
	}

	token, err := security.CreateToken(user.ID)
	if err != nil {
		helper.JSONResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	helper.JSONResponse(w, http.StatusOK, token)
}
