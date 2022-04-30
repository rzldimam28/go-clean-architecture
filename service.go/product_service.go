package service

import (
	"belajar-golang-clean-architecture/helper"
	"belajar-golang-clean-architecture/model/domain"
	"belajar-golang-clean-architecture/model/web"
	"belajar-golang-clean-architecture/repository"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
	Validator *validator.Validate
}

// list of method
func (service *ProductService) Create(request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)
	
	product := domain.Product{
		Name: request.Name,
	}

	product = service.ProductRepository.Save(product)
	return web.ProductResponse{
		Id: product.Id,
		Name: product.Name,
	}
}

func (service *ProductService) List() []web.ProductResponse {
	products := service.ProductRepository.List()

	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, web.ProductResponse{
			Id: product.Id,
			Name: product.Name,
		})
	}
	return productResponses
}