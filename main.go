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
	router.POST("/forum/user", routes.AddUser)
	router.GET("/forum/users", routes.ListUsers)
	router.GET("/forum/user/:userId", routes.GetUser)
	router.GET("/forum/:categoryId/topic", routes.ListTopics)
	router.POST("/forum/:categoryId/topic", routes.AddTopic)
	router.POST("/forum/:categoryId/topic/:postId/post", routes.AddPost)
	router.GET("/forum/:categoryId/topic/:postId/posts", routes.GetPosts)
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
