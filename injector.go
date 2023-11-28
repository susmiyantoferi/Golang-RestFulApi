//go:build wireinject
// +build wireinject

package main

import (
	"RestApi/app"
	"RestApi/controller"
	"RestApi/middleware"
	"RestApi/repository"
	"RestApi/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var bookSet = wire.NewSet(
	repository.NewBookRepository,
	wire.Bind(new(repository.BoookRepository), new(*repository.BookRepositoryImpl)),
	service.NewBookServiceImpl,
	wire.Bind(new(service.BookService), new(*service.BookServiceImpl)),
	controller.NewBookController,
	wire.Bind(new(controller.BookController), new(*controller.BookControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		bookSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
