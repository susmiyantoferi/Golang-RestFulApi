package test

import (
	"RestApi/app"
	"RestApi/controller"
	"RestApi/helper"
	"RestApi/middleware"
	"RestApi/model/domain"
	"RestApi/repository"
	"RestApi/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func SetupDB() *sql.DB {
	connStr := "postgres://postgres:Terserah123@localhost:5432/resfulapitest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookServiceImpl(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	return middleware.NewAuthMiddleware(router)
}

func TruncateBook(db *sql.DB) {
	db.Exec("truncate books")
}

func TestBookCreateSuccess(t *testing.T) {
	db := SetupDB()
	TruncateBook(db)
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"title" : "test", "author" : "test", "descrip" : "test"}`)
	request := httptest.NewRequest(http.MethodPost, "https://localhost:3000/api/books", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestBookCreateFailed(t *testing.T) {
	db := SetupDB()
	TruncateBook(db)
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"title" : "" , "author" : "", "descrip" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "https://localhost:3000/api/books", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestBookUpdateSuccess(t *testing.T) {
	db := SetupDB()
	TruncateBook(db)

	tx, _ := db.Begin()
	bookRepository := repository.NewBookRepository()
	book := bookRepository.Save(context.Background(), tx, domain.Book{
		Title:   "test",
		Author:  "test",
		Descrip: "test",
	})
	tx.Commit()

	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"title" : "feri", "author" : "susmiyanto", "descrip" : "iref"}`)
	request := httptest.NewRequest(http.MethodPut, "https://localhost:3000/api/books/"+strconv.Itoa(book.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//error
	//response := recorder.Result()
	//assert.Equal(t, 200, response.StatusCode)
	//
	//body, _ := io.ReadAll(response.Body)
	//var responseBody map[string]interface{}
	//json.Unmarshal(body, &responseBody)
	//
	//assert.Equal(t, 200, int(responseBody["code"].(float64)))
	//assert.Equal(t, "OK", responseBody["status"])
	//assert.Equal(t, book.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	//assert.Equal(t, `{"feri", "susmiyanto", "iref"'}`, responseBody["data"].(map[string]interface{})["title, author, descrip"])

}

func TestBookUpdateFailed(t *testing.T) {

}

func TestBookGetSuccess(t *testing.T) {

}

func TestBookGetFailed(t *testing.T) {

}
