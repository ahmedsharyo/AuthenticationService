package controllers

import (
	"github.com/ahmedsharyo/AuthenticationService/models"
	"github.com/ahmedsharyo/AuthenticationService/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) string {
	var credential models.User
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no correct data found in the payload"
	}
	isUserAuthenticated := service.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return service.GenerateToken(credential.Email, true)

	}
	return ""
}
