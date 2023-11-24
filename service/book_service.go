package service

import (
	"RestApi/model/web"
	"context"
)

// merupakan contract interface
type BookService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}
