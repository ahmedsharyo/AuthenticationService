package service

import (
	"fmt"

	Models "github.com/ahmedsharyo/AuthenticationService/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email string, password string) bool {

	user := Models.User{}
	if Models.GetUserByEmail(&user, email) != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	fmt.Print("true")
	return true
}
