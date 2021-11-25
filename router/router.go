package router

import (
	"net/http"

	"github.com/ahmedsharyo/AuthenticationService/controllers"
	"github.com/gin-gonic/gin"
)

func Routering(server *gin.Engine) {

	server.POST("/login", func(ctx *gin.Context) {
		token := controllers.Login(ctx)

		if token != "" {

			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})

		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	server.POST("/signup", func(ctx *gin.Context) {

		response := controllers.SignUp(ctx)
		ctx.JSON(http.StatusOK, gin.H{
			"response": response,
		})

	})

}
