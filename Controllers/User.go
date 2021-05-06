package Controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/ahmedsharyo/AuthenticationService/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var toID = map[string]Models.AuthorityLevel{
	"SystemAdmin":    Models.SystemAdmin,
	"EntityAdmin":    Models.EntityAdmin,
	"SiteAdmin":      Models.SiteAdmin,
	"SecurityManger": Models.SecurityManger,
}

var SecretKey = os.Getenv("SecretKey")

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//password hashing
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := Models.User{
		UserName:       data["username"],
		Email:          data["email"],
		Password:       string(password),
		AuthorityLevel: toID[data["authorityLevel"]],
	}

	Models.CreateUser(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user Models.User
	Models.GetUserByEmail(&user, data["email"])

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30 * 30).Unix(), //30 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// func User(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil
// 	})

// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		return c.JSON(fiber.Map{
// 			"message": "unauthenticated",
// 		})
// 	}

// 	claims := token.Claims.(*jwt.StandardClaims)

// 	var user models.User

// 	database.DB.Where("id = ?", claims.Issuer).First(&user)

// 	return c.JSON(user)
// }

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// //UpdateUser ... Update the User information
// func UpdateUser(c *gin.Context) {

// 	var user Models.User
// 	id := c.Params.ByName("id")
// 	err := Models.GetUserByID(&user, id)

// 	if err != nil {
// 		c.JSON(http.StatusNotFound, user)
// 	}
// 	c.BindJSON(&user)
// 	err = Models.UpdateUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// //DeleteUser ... Delete the User
// func DeleteUser(c *gin.Context) {
// 	var user Models.User
// 	id := c.Params.ByName("id")
// 	err := Models.DeleteUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
// 	}
// }
