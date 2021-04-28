package Controllers

import (
 "github.com/ahmedsharyo/SharyoPATv1/AuthenticationService/Models"
 "fmt"
 "net/http"
"github.com/gin-gonic/gin"
)


//GetsecurityMangers ... Get all securityMangers
func GetSecurityMangers(c *gin.Context) {

 var securityManger []Models.SecurityManger

 err := Models.GetAllSecurityMangers(&securityManger)

 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 }
 else {
  c.JSON(http.StatusOK, securityManger)
 }

}


//CreatesecurityManger ... Create securityManger
func CreateSecurityManger(c *gin.Context) {

 var securityManger Models.SecurityManger
 c.BindJSON(&securityManger)
 err := Models.CreateSecurityManger(&securityManger)

 if err != nil {
  fmt.Println(err.Error())
  c.AbortWithStatus(http.StatusNotFound)
 } 
 else {
  c.JSON(http.StatusOK, securityManger)
 }

}

//GetsecurityMangerByID ... Get the securityManger by id
func GetSecurityMangerByID(c *gin.Context) {
    
 id := c.Params.ByName("id")
 var securityManger Models.SecurityManger
 err := Models.GetSecurityMangerByID(&securityManger, id)

 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, securityManger)
 }

}

//UpdatesecurityManger ... Update the securityManger information
func UpdateSecurityManger(c *gin.Context) {

 var securityManger Models.SecurityManger
 id := c.Params.ByName("id")
 err := Models.GetSecurityMangerByID(&securityManger, id)
 
 if err != nil {
  c.JSON(http.StatusNotFound, securityManger)
 }
 c.BindJSON(&securityManger)
 err = Models.UpdateSecurityManger(&securityManger, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, securityManger)
 }
}

//DeletesecurityManger ... Delete the securityManger
func DeleteSecurityManger(c *gin.Context) {
 var securityManger Models.SecurityManger
 id := c.Params.ByName("id")
 err := Models.DeleteSecurityManger(&securityManger, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
 }
}