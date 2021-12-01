package controllers

import (
	"fmt"

	"github.com/ahmedsharyo/AuthenticationService/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) bool {

	user := models.User{}
	if models.GetUserByEmail(&user, email) != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	fmt.Print("true")
	return true
}
