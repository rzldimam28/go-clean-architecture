package controller

import (
	"belajar-golang-clean-architecture/helper"
	"belajar-golang-clean-architecture/model/web"
	"belajar-golang-clean-architecture/service.go"
	"encoding/json"
	"net/http"
)

type ProductController struct {
	ProductService *service.ProductService
}

// list of method
func (controller *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	// read from request body
	productCreateRequest := web.ProductCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&productCreateRequest)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.Create(productCreateRequest)
	webResponse := web.CreateWebResponse(200, "Success Create Product", productResponse)

	// write to request body
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *ProductController) List(w http.ResponseWriter, r *http.Request) {
	products := controller.ProductService.List()
	w.Header().Add("Content-Type", "application/json")
	webResponse := web.CreateWebResponse(200, "Success Get All Products", products)
	// write to request body
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
	
}