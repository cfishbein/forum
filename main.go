package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cfishbein/forum/internal/db"
	"github.com/cfishbein/forum/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := loadConf()
	log.Println("Loading database")
	db.InitDB(conf.DatabasePath)
	defer db.Close()

	routes.RegisterCategories()

	router := gin.New()
	router.POST("/forum/users", routes.AddUser)
	router.GET("/forum/users", routes.ListUsers)
	router.GET("/forum/users/:userId", routes.GetUser)
	router.GET("/forum/categories/:categoryId/threads", routes.ListThreads)
	router.POST("/forum/categories/:categoryId/threads", routes.AddThread)
	router.POST("/forum/categories/:categoryId/threads/:threadId/posts", routes.AddPost)
	router.GET("/forum/categories/:categoryId/threads/:threadId/posts", routes.ListPosts)
	router.Run(conf.Port)
}

type configuration struct {
	Port         string `json:"port"`
	DatabasePath string `json:"databasePath"`
}

func loadConf() configuration {
	file, err := os.Open(os.Getenv("FORUM_CFG"))
	if err != nil {
		panic("Failed to open configuration file specified by FORUM_CFG env")
	}
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
