package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func JSONResponseWithError(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
