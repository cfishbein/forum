package db

import (
	"database/sql"
	"os"

	"github.com/cfishbein/forum/internal/model"
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

// ListCategories lists all available Categories
func ListCategories() ([]model.Category, error) {
	rows, err := db.Query("SELECT id, name, desc FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make([]model.Category, 0)
	for rows.Next() {
		var id int
		var name string
		var desc string
		err = rows.Scan(&id, &name, &desc)
		if err != nil {
			return nil, err
		}
		category := model.Category{ID: id, Name: name, Desc: desc}
		categories = append(categories, category)
	}
	return categories, nil
}

// AddUser adds the given user to the database
func AddUser(u model.User) error {
	err := exec("INSERT INTO user(name) VALUES(?)", func(stmt *sql.Stmt) error {
		_, e := stmt.Exec(u.Name)
		return e
	})
	return err
}

// GetUser gets a user with the given ID
func GetUser(id int) (*model.User, error) {
	rows, err := db.Query("SELECT name FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	var name string
	err = rows.Scan(&name)
	if err != nil {
		return nil, err
	}
	user := &model.User{ID: id, Name: name}
	return user, nil
}

// ListUsers lists all users in the database
func ListUsers() ([]model.User, error) {
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]model.User, 0)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		user := model.User{ID: id, Name: name}
		users = append(users, user)
	}
	return users, nil
}

// GetCategory gets the Category for the given ID
func GetCategory(id int) (*model.Category, error) {
	rows, err := db.Query("SELECT name, desc FROM category WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	var name string
	var desc string
	err = rows.Scan(&name, &desc)
	if err != nil {
		return nil, err
	}
	category := &model.Category{ID: id, Name: name, Desc: desc}
	return category, nil
}

// AddThread adds a new Thread, setting the thread ID of t
func AddThread(c model.Category, t *model.Thread) error {
	err := exec("INSERT INTO thread(title, author_id, category_id) VALUES(?, ?, ?)", func(stmt *sql.Stmt) error {
		result, e := stmt.Exec(t.Title, t.Author.ID, c.ID)
		if e != nil {
			return e
		}
		id, e := result.LastInsertId()
		t.ID = int(id)
		return e
	})
	return err
}

// ListThreads lists all threads for the given category ID
func ListThreads(id int) ([]model.Thread, error) {
	rows, err := db.Query("SELECT id, title, author_id FROM thread WHERE category_id = ?")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	threads := make([]model.Thread, 0)
	for rows.Next() {
		var id int
		var title string
		var authorID int
		err = rows.Scan(&id, &title, &authorID)
		if err != nil {
			return nil, err
		}
		author, err := GetUser(authorID)
		if err != nil {
			return nil, err
		}
		thread := model.Thread{ID: id, Title: title, Author: *author}
		threads = append(threads, thread)
	}
	return threads, nil
}

// AddPost adds the given Post to the database
func AddPost(threadID int, p model.Post) error {
	err := exec("INSERT INTO post(thread_id, author_id, content) VALUES(?, ?, ?)", func(stmt *sql.Stmt) error {
		_, e := stmt.Exec(threadID, p.Author.ID, p.Content)
		return e
	})
	return err
}

// GetPosts get all posts for the given Thread ID
func GetPosts(id int) ([]model.Post, error) {
	rows, err := db.Query("SELECT p.id as id, a.id as author_id, a.name, p.content FROM post p, user a WHERE thread_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]model.Post, 0)
	for rows.Next() {
		var postID int
		var authorID int
		var name string
		var content string

		err = rows.Scan(&postID, &authorID, &name, &content)
		if err != nil {
			return nil, err
		}
		author := model.User{ID: authorID, Name: name}
		post := model.Post{ID: postID, Author: author, Content: content}
		posts = append(posts, post)
	}
	return posts, nil
}

// exec convenience function to execute the given SQL string and uses the given fn to set execution parameters
func exec(sql string, fn func(*sql.Stmt) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = fn(stmt)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
