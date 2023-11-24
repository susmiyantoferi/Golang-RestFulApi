package service

import (
	"RestApi/exception"
	"RestApi/helper"
	"RestApi/model/domain"
	"RestApi/model/web"
	"RestApi/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BoookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repository.BoookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	book := domain.Book{
		Title:   request.Title,
		Author:  request.Author,
		Descrip: request.Descrip,
	}

	book = service.BookRepository.Save(ctx, tx, book)
	return helper.ToBookResponse(book)

}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.Title = request.Title
	book.Author = request.Author
	book.Descrip = request.Descrip

	book = service.BookRepository.Update(ctx, tx, book)
	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)

}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)

	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, helper.ToBookResponse(book))
	}

	return helper.ToBookResponses(books)
}
