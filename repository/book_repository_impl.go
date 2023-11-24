package repository

import (
	"RestApi/helper"
	"RestApi/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BoookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	//SQL := "insert into books (title, author, descrip) values ($1, $2, $3)"

	SQL := `insert into books (title, author, descrip) values ($1, $2, $3) RETURNING id`
	result, err := tx.ExecContext(ctx, SQL, book.Title, book.Author, book.Descrip)
	helper.PanicIfError(err)

	id, err := result.RowsAffected()
	helper.PanicIfError(err)

	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	//SQL := "update books set title, author, descrip = ($1, $2, $3) where id = ($4)"
	SQL := `update books set (title, author, descrip) = ($1, $2, $3) where id = ($4) RETURNING id`
	_, err := tx.ExecContext(ctx, SQL, book.Title, book.Author, book.Descrip, book.Id)

	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	//SQL := "delete from books where id = ($1)"
	SQL := `delete from books where id = ($1)`
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.PanicIfError(err)
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	SQL := "select id, title, author, descrip from books where id = ($1)"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Descrip)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("boooks id not found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "select id, title, author, descrip from books"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Descrip)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}
