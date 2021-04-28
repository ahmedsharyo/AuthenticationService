package Controllers

import (
 "github.com/ahmedsharyo/AuthenticationService/Models"
 "fmt"
 "net/http"
"github.com/gin-gonic/gin"
)

//GetAdmins ... Get all Admins
func GetAdmins(c *gin.Context) {

 var admin []Models.Admin

 err := Models.GetAllAdmins(&admin)

 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 }else {
  c.JSON(http.StatusOK, admin)
 }

}


//CreateAdmin ... Create Admin
func CreateAdmin(c *gin.Context) {

 var admin Models.Admin
 c.BindJSON(&admin)
 err := Models.CreateAdmin(&admin)

 if err != nil {
  fmt.Println(err.Error())
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, admin)
 }

}

//GetAdminByID ... Get the Admin by id
func GetAdminByID(c *gin.Context) {
    
 id := c.Params.ByName("id")
 var admin Models.Admin
 err := Models.GetAdminByID(&admin, id)

 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, admin)
 }

}

//UpdateAdmin ... Update the Admin information
func UpdateAdmin(c *gin.Context) {

 var admin Models.Admin
 id := c.Params.ByName("id")
 err := Models.GetAdminByID(&admin, id)
 
 if err != nil {
  c.JSON(http.StatusNotFound, admin)
 }
 c.BindJSON(&admin)
 err = Models.UpdateAdmin(&admin, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, admin)
 }
}

//DeleteAdmin ... Delete the Admin
func DeleteAdmin(c *gin.Context) {
 var admin Models.Admin
 id := c.Params.ByName("id")
 err := Models.DeleteAdmin(&admin, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
 }
}