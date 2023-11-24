package controller

import (
	"RestApi/helper"
	"RestApi/model/web"
	"RestApi/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFormRequestBody(request, &bookCreateRequest)

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFormRequestBody(request, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) FindALl(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookResponses := controller.BookService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
