package model

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
