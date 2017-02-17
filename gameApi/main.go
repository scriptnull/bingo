package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"

	gin "gopkg.in/gin-gonic/gin.v1"
)

type gameAPIConfig struct {
	Port       int    `json:"port"`
	DbHost     string `json:"dbHost"`
	DbName     string `json:"dbName"`
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
}

type appConfig struct {
	GameAPI gameAPIConfig `json:"gameApi"`
}

func main() {

	// read app config file
	rawConfig, err := ioutil.ReadFile("../app.config.json")
	if err != nil {
		log.Fatalf("Unable to load configuration: %v", err)
	}

	// load app config
	var app = appConfig{}
	err = json.Unmarshal(rawConfig, &app)
	if err != nil {
		log.Fatalf("Unable to parse app.config.json: %v", err)
	}

	// connect to db
	var connectionString = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		app.GameAPI.DbHost,
		app.GameAPI.DbUser,
		app.GameAPI.DbName,
		app.GameAPI.DbPassword,
	)
	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	// Start gin
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
	})

	router.Run(string(app.GameAPI.Port))
}
