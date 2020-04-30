package model

import (
	"errors"
	"strings"
)

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

// NewPost creates a new post with the given user and content
func NewPost(author User, content string) (*Post, error) {
	if len(strings.TrimSpace(content)) < 1 {
		return nil, errors.New("Post must contain content")
	}
	p := new(Post)
	p.Author = author
	p.Content = content
	return p, nil
}
