package controllers

// import (
// 	"github.com/ahmedsharyo/AuthenticationService/models"
// 	"github.com/ahmedsharyo/AuthenticationService/service"
// 	"github.com/gin-gonic/gin"
// )

// func SignUp(ctx *gin.Context) string {

// 	var data models.User
// 	err := ctx.ShouldBind(&data)
// 	if err != nil {
// 		return "no correct data found in the payload"
// 	}
// 	if err := service.SignUp(data); err != nil {
// 		return "db err"
// 	}

// 	return "user added"
// }
