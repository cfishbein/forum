package routes

import (
	"net/http"
	"strconv"

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
