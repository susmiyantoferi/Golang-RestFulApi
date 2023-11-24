package helper

import (
	"RestApi/model/domain"
	"RestApi/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:      book.Id,
		Title:   book.Title,
		Author:  book.Author,
		Descrip: book.Descrip,
	}
}

// konversi response
func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}
	return bookResponses

}
