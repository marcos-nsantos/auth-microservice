package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/e-commerce/auth-service/handler"
	"github.com/marcos-nsantos/e-commerce/auth-service/middleware"
)

func AuthHandleRequestsV1(v1 chi.Router) {
	v1.Use(middleware.Authenticate)

	v1.Get("/{id}", handler.FindUserByID)
	v1.Put("/{id}", handler.UpdateUser)
	v1.Patch("/changePassword/{id}", handler.UpdateUserPassword)
	v1.Delete("/{id}", handler.DeleteUser)
}

func NoAuthHandleRequestsV1(v1 chi.Router) {
	v1.Post("/", handler.CreateUser)
	v1.Post("/login", handler.Login)
}
