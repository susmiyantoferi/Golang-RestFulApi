package app

import (
	"RestApi/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	connStr := "postgres://postgres:Terserah123@localhost:5432/resfulapigo?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
