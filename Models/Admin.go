package Models

import (
 "github.com/ahmedsharyo/AuthenticationService/Config"
 "fmt"
_ "github.com/go-sql-driver/mysql"
)

//GetAllAdmins Fetch all Admin data
func GetAllAdmins(admin *[]Admin) (err error) {
 if err = Config.DB.Find(admin).Error; err != nil {
  return err
 }
 return nil
}

//CreateAdmin ... Insert New data
func CreateAdmin(admin *Admin) (err error) {
 if err = Config.DB.Create(admin).Error; err != nil {
  return err
 }
 return nil
}


//GetAdminByID ... Fetch only one Admin by Id
func GetAdminByID(admin *admin, id string) (err error) {
 if err = Config.DB.Where("id = ?", id).First(admin).Error; err != nil {
  return err
 }
 return nil
}

//UpdateAdmin ... Update Admin
func UpdateAdmin(Admin *admin, id string) (err error) {
 fmt.Println(admin)
 Config.DB.Save(admin)
 return nil
}

//DeleteAdmin ... Delete Admin
func DeleteAdmin(admin *Admin, id string) (err error) {
 Config.DB.Where("id = ?", id).Delete(admin)
 return nil
}
