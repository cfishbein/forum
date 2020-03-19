package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type topic struct {
	Title  string `json:"title"`
	Author user   `json:"author"`
	Posts  []post `json:"posts"`
}

type post struct {
	Author  user   `json:"author"`
	Content string `json:"content"`
}

func openDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	return db
}
