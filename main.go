package main

import (
	"fmt"

	"github.com/ahmedsharyo/AuthenticationService/Config"
	"github.com/ahmedsharyo/AuthenticationService/Models"
	"github.com/ahmedsharyo/AuthenticationService/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	Routes.Setup(app)
	//Tests.SetApp(app)

	app.Listen(":8080")

}
