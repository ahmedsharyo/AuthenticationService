package Models

import (
	"fmt"

	"github.com/ahmedsharyo/AuthenticationService/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all User data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one User by Id
func GetUserByEmail(user *User, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update User
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete User
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
