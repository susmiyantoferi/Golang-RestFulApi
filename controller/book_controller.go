package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindALl(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
