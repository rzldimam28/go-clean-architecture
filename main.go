package main

import (
	"belajar-golang-clean-architecture/cache"
	"belajar-golang-clean-architecture/controller"
	"belajar-golang-clean-architecture/helper"
	"belajar-golang-clean-architecture/repository"
	"belajar-golang-clean-architecture/service.go"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func PanicRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Fprintln(w, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func InitRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func main() {
	
	// create db
	db, err := sql.Open("mysql", "root:leesrcyng__@tcp(localhost:3306)/clean_architecture_go")
	helper.PanicIfError(err)
	validator := validator.New()

	// create redis
	client := InitRedisClient()
	redisC := cache.NewRedisCache(client, 1)

	// repo, service and controller
	productRepository := repository.ProductRepository{DB: db}
	productService := service.ProductService{ProductRepository: &productRepository, Validator: validator, Cache: redisC}
	productController := controller.ProductController{ProductService: &productService}

	// initiate http
	r := mux.NewRouter()

	// middleware for displaying error msg
	// r.Use(PanicRecover)

	// route
	r.HandleFunc("/products", productController.Create).Methods("POST")
	r.HandleFunc("/products", productController.List).Methods("GET")

	log.Println("Server is running")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}