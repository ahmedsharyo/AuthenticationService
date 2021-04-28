package main
import (
 "github.com/ahmedsharyo/AuthenticationService/Config"
 "github.com/ahmedsharyo/AuthenticationService/Models"
 "github.com/ahmedsharyo/AuthenticationService/Routes"
 "fmt"
"github.com/jinzhu/gorm"
)

var err error
func main() {

 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

if err != nil {
  fmt.Println("Status:", err)
 }

defer Config.DB.Close()

Config.DB.AutoMigrate(&Models.User{})
Config.DB.AutoMigrate(&Models.Admin{})
Config.DB.AutoMigrate(&Models.SecurityManger{})


r := Routes.SetupRouter()
 //running
 r.Run()
}
