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

// Category is a forum category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Thread is a forum thread
type Thread struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author User   `json:"author"`
	Posts  []Post `json:"posts"`
}

// NewThread creates a new thread with the given title and user
func NewThread(title string, author User) (*Thread, error) {
	if len(strings.TrimSpace(title)) < 1 {
		return nil, errors.New("Title must not be empty")
	}
	t := &Thread{}
	t.Author = author
	t.Title = title
	return t, nil
}

// Post is a post within a Thread
type Post struct {
	ID      int    `json:"id"`
	Author  User   `json:"author"`
	Content string `json:"content"`
}

// NewPost creates a new post with the given content and user
func NewPost(content string, author User) (*Post, error) {
	if len(strings.TrimSpace(content)) < 1 {
		return nil, errors.New("Post must contain content")
	}
	p := new(Post)
	p.Author = author
	p.Content = content
	return p, nil
}
