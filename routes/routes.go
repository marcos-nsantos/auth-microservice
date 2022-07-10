package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func HandleRequests() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/v1/auth/users", AuthHandleRequestsV1)
	r.Route("/v1/users", NoAuthHandleRequestsV1)

	return r
}
