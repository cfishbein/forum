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

// NewTopic creates a new topic with the given title and user
func NewTopic(title string, author User) (*Topic, error) {
	if len(strings.TrimSpace(title)) < 1 {
		return nil, errors.New("Title must not be empty")
	}
	t := &Topic{}
	t.Author = author
	t.Title = title
	return t, nil
}

// Post is a post within a topic
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
