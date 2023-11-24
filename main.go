package main

import (
	"RestApi/app"
	"RestApi/controller"
	"RestApi/helper"
	"RestApi/middleware"
	"RestApi/repository"
	"RestApi/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookServiceImpl(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
