package main

import (
	"fmt"

	Config "github.com/ahmedsharyo/AuthenticationService/config"
	Models "github.com/ahmedsharyo/AuthenticationService/models"
	"github.com/ahmedsharyo/AuthenticationService/router"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	server := gin.New()

	router.Routering(server)

	port := "8080"
	server.Run(":" + port)
}
