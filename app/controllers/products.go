package controllers

import (
	"FruitSale/app/Entities"
	"FruitSale/app/models"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

/*
 Request Patterns for Products
*/

var (
	listProducts  = regexp.MustCompile(`^\/api\/v1\/products[\/]*$`)
	singleProduct = regexp.MustCompile(`^\/api\/v1\/products\/\d+$`)
)

/*
Request Handler for Products
*/
func HandleProductRequests(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	switch {
	case request.Method == http.MethodPut && singleProduct.MatchString(request.URL.Path):
		Update(writer, request)
		return
	case request.Method == http.MethodGet && singleProduct.MatchString(request.URL.Path):
		ListProduct(writer, request)
		return
	case request.Method == http.MethodGet && listProducts.MatchString(request.URL.Path):
		ListProducts(writer, request)
		return
	default:
		errorResponse(writer, "Invalid Url - check api doc", http.StatusNotFound)
		return
	}
}

/*
ListProducts

	List all products from Memory storage
	By Default only display 10 products per page
	Accepts two params offset & limit
	offset -> Starting nmber for the products, by default first page will start from 1
	limit -> Total No of products to be displayed in single page, default value 10
*/
func ListProducts(w http.ResponseWriter, r *http.Request) {
	if verifyHeader(w, r) {
		offset, err1 := strconv.Atoi(r.URL.Query().Get("offset"))
		if err1 != nil {
			offset = 0
		} else {
			offset--
		}
		limit, err2 := strconv.Atoi(r.URL.Query().Get("limit"))
		if err2 != nil {
			limit = 10
		}
		products, _ := models.ListAllProducts(offset, limit)
		response, _ := json.Marshal(products)
		w.Write(response)
	}
}

/*
ListProduct

	List details of the single product
*/
func ListProduct(w http.ResponseWriter, r *http.Request) {
	if verifyHeader(w, r) {
		productId := getProductId(w, r)
		product, error := models.ListProduct(productId)
		if error != nil {
			errorResponse(w, string(error.Error()), http.StatusNotFound)
			return
		}
		response, _ := json.Marshal(product)
		w.Write(response)
	}
}

/*
Update

	Updates the title of the given product
*/
func Update(w http.ResponseWriter, r *http.Request) {
	if verifyHeader(w, r) && isAuthorized(w, r) {
		productId := getProductId(w, r)
		var e Entities.Product
		var unmarshalErr *json.UnmarshalTypeError
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&e)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			}
			return
		}
		e.ID = productId
		success, failed := models.UpdateProduct(e)
		if failed != nil {
			errorResponse(w, failed.Error(), http.StatusConflict)
			return
		}
		newFsConfigBytes, _ := json.Marshal(success)
		w.Write(newFsConfigBytes)
	}
}

/*
Helper functions
*/
func verifyHeader(w http.ResponseWriter, r *http.Request) bool {
	return true
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok || !isValidOwner(username, password) {
		errorResponse(w, "Login to update data / Invalid Credentials", http.StatusUnauthorized)
		return false
	}
	return true
}

func isValidOwner(user string, password string) bool {
	return user == "admin" && password == "admin"
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func getProductId(w http.ResponseWriter, r *http.Request) int {
	regex := regexp.MustCompile(`(\d+$)`)
	productId, err := strconv.Atoi(regex.FindStringSubmatch(r.URL.Path)[1])
	if err != nil {
		errorResponse(w, "Product not found", http.StatusNotFound)
		return 0
	}
	return productId
}
