package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"demo-go-basic-backend/models"
)

func respondWithJSON(w http.ResponseWriter, code int, response models.JSON) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func SuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	var response = models.JSON{Code: code, Message: "OK", Data: data, Error: false}

	respondWithJSON(w, code, response)
}

func ErrorResponse(w http.ResponseWriter, code int, errMessage string) {
	var response = models.JSON{Code: code, Error: true, ErrorMessage: errMessage}

	respondWithJSON(w, code, response)
}
