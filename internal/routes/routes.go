package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cfishbein/forum/internal/db"
	"github.com/cfishbein/forum/internal/model"
	"github.com/gin-gonic/gin"
)

// AddUser attempts to add a new User
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	user := model.User{Name: name}
	err := db.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{})
	}
}

// GetUser attempts to get an existing User
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid user ID"})
		return
	}
	user, err := db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// ListUsers lists all users
func ListUsers(c *gin.Context) {
	users, err := db.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

// GetPosts gets all posts for a topic ID in the path param
func GetPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid post ID"})
		return
	}
	posts, err := db.GetPosts(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	}
}

// AddPost adds a post to the database
func AddPost(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic ID"})
		return
	}
	userID, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	content := c.PostForm("content")
	if len(strings.TrimSpace(content)) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post must contain content"})
		return
	}

	author, err := db.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User ID not found"})
		return
	}

	// TODO FK's not being enforce in sqlite3, so topic ID isn't validated
	post := new(model.Post)
	post.Author = *author
	post.Content = content
	err = db.AddPost(topicID, *post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
