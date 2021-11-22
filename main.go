package main

import (
	"net/http"

	"github.com/ahmedsharyo/AuthenticationService/controllers"
	"github.com/ahmedsharyo/AuthenticationService/service"
	"github.com/gin-gonic/gin"
)

//var err error

func main() {

	// Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }

	// defer Config.DB.Close()

	// Config.DB.AutoMigrate(&Models.User{})

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	port := "8080"
	server.Run(":" + port)
}
