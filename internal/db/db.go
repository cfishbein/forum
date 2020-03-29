package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3" // TODO check how to import (example used this)
)

// don't expect this app to really ever be used by more than
// just myself, so global db pattern is fine for now
var db *sql.DB

// InitDB initializes the database at the given path
func InitDB(path string) {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if os.IsNotExist(err) {
		panic("Configured database path does not exist")
	}
	handle, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	db = handle
}

// Close closes the database
func Close() {
	db.Close()
}

// User is a user account
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Topic is a forum topic
type Topic struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author User   `json:"author"`
	Posts  []Post `json:"posts"`
}

// Post is a post within a topic
type Post struct {
	ID      int    `json:"id"`
	Author  User   `json:"author"`
	Content string `json:"content"`
}

// InsertUser inserts the given user to the database
func InsertUser(u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO user(name) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// ListUsers lists all users in the database
func ListUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		user := User{ID: id, Name: name}
		users = append(users, user)
	}
	return users, nil
}

// GetPosts get all posts for the given Topic ID
func GetPosts(id int) ([]Post, error) {
	rows, err := db.Query("SELECT p.id as id, a.id as author_id, a.name, p.content FROM post p, user a WHERE topic_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]Post, 0)
	for rows.Next() {
		var postID int
		var authorID int
		var name string
		var content string

		err = rows.Scan(&postID, &authorID, &name, &content)
		if err != nil {
			return nil, err
		}
		author := User{ID: authorID, Name: name}
		post := Post{ID: postID, Author: author, Content: content}
		posts = append(posts, post)
	}
	return posts, nil
}
