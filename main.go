package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"demo-go-basic-backend/handlers"
)

func main() {
	r := mux.NewRouter()

	// ENDPOINT FOR /PRODUCTS
	r.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", handlers.GetSingleProduct).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods(http.MethodDelete)

	r.Use(handlers.LoggingMiddleware)

	// handling 404
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// handling wrong methods
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedHandler)

	// ListenAndServe listens on the TCP network address addr and then calls Serve with handlers to handle requests on incoming connections.
	fmt.Printf("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
