package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // TODO check how to import (example used this)
)

// User is a user account
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Topic is a forum topic
type Topic struct {
	Title  string `json:"title"`
	Author User   `json:"author"`
	Posts  []Post `json:"posts"`
}

// Post is a post within a topic
type Post struct {
	Author  User   `json:"author"`
	Content string `json:"content"`
}

// OpenDB opens the database file
func OpenDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	return db
}
