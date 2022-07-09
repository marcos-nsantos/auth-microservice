package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/handler"
)

func HandleRequestsV1(v1 chi.Router) {
	v1.Post("/users", handler.CreateUser)
}
