package security

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(userId uint) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign and get the complete encoded token as a string
	return token.SignedString(key)
}

func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)
	if len(tokenString) == 0 {
		return errors.New("token not found")
	}

	token, err := jwt.Parse(tokenString, validateKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return nil
	}

	return jwt.ErrSignatureInvalid
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(token) == 0 {
		return ""
	}

	if token[0:7] == "Bearer " {
		return token[7:]
	}

	return ""
}

func validateKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return key, nil
}
