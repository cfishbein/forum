package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := loadConf()
	router := gin.New()
	router.GET("/users", listUsers)
	router.Run(conf.Port)
}

type configuration struct {
	Port string `json:"port"`
}

func loadConf() configuration {
	file, err := os.Open("conf.json")
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
