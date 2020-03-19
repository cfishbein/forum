package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := loadConf()
	log.Println("Loading database")
	db := openDB(conf.DatabasePath)
	defer db.Close()
	router := gin.New()
	router.GET("/users", listUsers)
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
