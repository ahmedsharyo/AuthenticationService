package service

import (
	"github.com/ahmedsharyo/AuthenticationService/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(data models.User) error {

	//password hashing
	pass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	data.Password = string(pass)

	return models.CreateUser(&data)
}
