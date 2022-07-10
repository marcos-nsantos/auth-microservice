package middleware

import (
	"github.com/marcos-nsantos/e-commerce/auth-service/helper"
	"github.com/marcos-nsantos/e-commerce/auth-service/security"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			helper.JSONResponseWithError(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
