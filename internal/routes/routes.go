package routes

import (
	"net/http"

	"github.com/cfishbein/forum/internal/db"
	"github.com/gin-gonic/gin"
)

var users [2]db.User = [2]db.User{db.User{ID: 1, Name: "user1"}, db.User{ID: 2, Name: "user2"}}

// ListUsers lists all users
func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
