package routes

import (
	"net/http"

	"github.com/cfishbein/forum/internal/db"
	"github.com/gin-gonic/gin"
)

// InsertUser attempts to insert a new User
func InsertUser(c *gin.Context) {
	name := c.PostForm("name")
	user := db.User{Name: name}
	err := db.InsertUser(user)
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
