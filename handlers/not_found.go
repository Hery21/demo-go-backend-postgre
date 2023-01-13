package handlers

import (
	"demo-go-basic-backend/helpers"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ErrorResponse(w, http.StatusNotFound, "404 not found!")
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed!")
}
