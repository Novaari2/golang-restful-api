package app

import (
	"database/sql"
	"go-restful/helper"
	"time"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/golang-restful-api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
