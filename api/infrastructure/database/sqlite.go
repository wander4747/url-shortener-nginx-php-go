package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type sqlite struct {
	DB *sql.DB
}

func NewSqlite() (*sqlite, error) {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatal(err)
		return &sqlite{}, nil
	}

	return &sqlite{DB: db}, nil
}
