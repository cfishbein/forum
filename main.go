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
	router.POST("/user", routes.AddUser)
	router.GET("/users", routes.ListUsers)
	router.GET("/user/:id", routes.GetUser)
	router.POST("/category/:categoryId/topic", routes.AddTopic)
	router.POST("/category/:categoryId/topic/:id/post", routes.AddPost)
	router.GET("/category/:categoryId/topic/:id/posts", routes.GetPosts)
	router.Run(conf.Port)
}

type configuration struct {
	Port         string `json:"port"`
	DatabasePath string `json:"databasePath"`
}

func loadConf() configuration {
	// TODO conf should really be from an environment variable
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
