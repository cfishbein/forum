package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users [2]string = [2]string{"user-1", "user-2"}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
