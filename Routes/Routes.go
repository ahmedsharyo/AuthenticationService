package Routes

import (
 "github.com/ahmedsharyo/AuthenticationService/Controllers"
"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

 r := gin.Default()

 grp1 := r.Group("/admin-api")
 {
  grp1.GET("admin", Controllers.GetAdmins)
  grp1.POST("admin", Controllers.CreateAdmin)
  grp1.GET("admin/:id", Controllers.GetAdminByID)
  grp1.PUT("admin/:id", Controllers.UpdateAdmin)
  grp1.DELETE("admin/:id", Controllers.DeleteAdmin)
 }
 
 grp2 := r.Group("/securityManger-api")
 {
  grp2.GET("securityManger", Controllers.GetSecurityMangers)
  grp2.POST("securityManger", Controllers.CreateSecurityManger)
  grp2.GET("securityManger/:id", Controllers.GetSecurityMangerByID)
  grp2.PUT("securityManger/:id", Controllers.UpdateSecurityManger)
  grp2.DELETE("securityManger/:id", Controllers.DeleteSecurityManger)
 }

 return r
}
