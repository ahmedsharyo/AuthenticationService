package Models

import (
 "github.com/ahmedsharyo/SharyoPATv1/AuthenticationService/Config"
 "fmt"
_ "github.com/go-sql-driver/mysql"
)

//GetAllSecurityMangers Fetch all SecurityManger data
func GetAllSecurityMangers(securityManger *[]SecurityManger) (err error) {
 if err = Config.DB.Find(securityManger).Error; err != nil {
  return err
 }
 return nil
}

//CreateSecurityManger ... Insert New data
func CreateSecurityManger(securityManger *SecurityManger) (err error) {
 if err = Config.DB.Create(securityManger).Error; err != nil {
  return err
 }
 return nil
}


//GetSecurityMangerByID ... Fetch only one SecurityManger by Id
func GetSecurityMangerByID(securityManger *SecurityManger, id string) (err error) {
 if err = Config.DB.Where("id = ?", id).First(securityManger).Error; err != nil {
  return err
 }
 return nil
}

//UpdateSecurityManger ... Update SecurityManger
func UpdateSecurityManger(securityManger *SecurityManger, id string) (err error) {
 fmt.Println(securityManger)
 Config.DB.Save(securityManger)
 return nil
}

//DeleteSecurityManger ... Delete SecurityManger
func DeleteSecurityManger(securityManger *SecurityManger, id string) (err error) {
 Config.DB.Where("id = ?", id).Delete(securityManger)
 return nil
}
