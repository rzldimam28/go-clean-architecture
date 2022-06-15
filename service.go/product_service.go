package service

import (
	"belajar-golang-clean-architecture/cache"
	"belajar-golang-clean-architecture/helper"
	"belajar-golang-clean-architecture/model/domain"
	"belajar-golang-clean-architecture/model/web"
	"belajar-golang-clean-architecture/repository"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
	Validator *validator.Validate
	Cache cache.ProductCache
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
	keys := "products"
	productsC, _ := service.Cache.Get(context.Background(), keys)
	fmt.Println("productsC", productsC)

	var productRs []web.ProductResponse
	if productsC == nil {
		fmt.Println("masuk")

		products := service.ProductRepository.List()
		
		var productResponses []web.ProductResponse
		for _, product := range products {
			productResponses = append(productResponses, web.ProductResponse{
				Id: product.Id,
				Name: product.Name,
			})
		}

		service.Cache.Set(context.Background(), keys, productResponses)
		productRs = productResponses
		return productRs
	}

	for _, value := range productsC {
		productRs = append(productRs, web.ProductResponse{
			Id: value.Id,
			Name: value.Name,
		})
	}

	return productRs
}