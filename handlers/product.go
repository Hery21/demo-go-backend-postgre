package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"demo-go-basic-backend/helpers"
	"demo-go-basic-backend/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	p, err := models.GetProducts()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.SuccessResponse(w, http.StatusOK, p)
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	// get params
	params := mux.Vars(r)

	// convert from string to int
	id, err := strconv.Atoi(params["id"])

	// if error converting from string to int, return error "Invalid Id Format"
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format!")
		return
	}

	p := models.Product{ID: id}

	if err := p.GetSingleProduct(); err != nil {
		switch err {
		case sql.ErrNoRows:
			helpers.ErrorResponse(w, http.StatusBadRequest, "Product not found")
		default:
			helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	helpers.SuccessResponse(w, http.StatusOK, p)

	// search id in products slice
	//for _, product := range models.Products {
	//	if product.ID == id {
	//		helpers.SuccessResponse(w, http.StatusOK, product)
	//		return
	//	}
	//}
	//
	//// if Id not found, return "Id not found!"
	//helpers.ErrorResponse(w, http.StatusBadRequest, "ID not found!")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product

	// decode json from request
	decoder := json.NewDecoder(r.Body)

	// check if valid json
	if err := decoder.Decode(&p); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Manually validate input
	if p.Name == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing name field!")
		return
	}

	if p.Description == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing description field!")
		return
	}

	if p.Quantity == 0 {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing quantity field!")
		return
	}

	if err := p.CreateProduct(); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// random Id
	//p.ID = rand.Intn(100) + 100
	//models.Products = append(models.Products, p)
	helpers.SuccessResponse(w, http.StatusCreated, p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	// if error converting from string to int, return error "Invalid Id Format"
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format!")
		return
	}

	var p models.Product

	// decode json from request
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.ID = id

	// Manually validate input
	if p.Name == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing name field!")
		return
	}

	if p.Description == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing description field!")
		return
	}

	if p.Quantity == 0 {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Missing quantity field!")
		return
	}

	if err := p.UpdateProduct(); err != nil {
		switch err {
		case sql.ErrNoRows:
			helpers.ErrorResponse(w, http.StatusBadRequest, "Product not found")
		default:
			helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	helpers.SuccessResponse(w, http.StatusOK, p)

	//for index, product := range models.Products {
	//	if product.ID == id {
	//		product := &models.Products[index]
	//		_ = json.NewDecoder(r.Body).Decode(&product)
	//		helpers.SuccessResponse(w, http.StatusOK, product)
	//		return
	//	}
	//}
	//
	//// if Id not found, return "Id not found!"
	//helpers.ErrorResponse(w, http.StatusBadRequest, "ID not found!")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	p := models.Product{ID: id}

	if err := p.DeleteProduct(); err != nil {
		switch err {
		case sql.ErrNoRows:
			helpers.ErrorResponse(w, http.StatusBadRequest, "Product not found")
		default:
			helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	helpers.SuccessResponse(w, http.StatusOK, p)

	// if error converting from string to int, return error "Invalid Id Format"
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format!")
		return
	}

	for index, product := range models.Products {
		if product.ID == id {
			models.Products = append(models.Products[:index], models.Products[index+1:]...)
			helpers.SuccessResponse(w, http.StatusOK, product)
			return
		}
	}

	// if Id not found, return "Id not found!"
	helpers.ErrorResponse(w, http.StatusBadRequest, "ID not found!")
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
