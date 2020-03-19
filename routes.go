package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users [2]user = [2]user{user{ID: 1, Name: "user1"}, user{ID: 2, Name: "user2"}}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
