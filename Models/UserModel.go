package Models
// User Struct


type authorityLevel int

const ( 
    SystemAdmin authorityLevel = iota
	SiteAdmin
    EntityAdmin
    SecurityManger
)

type User struct {
    Id      uint   `json:"id"`
    UserName    string `json:"username"`
    Email   string `json:"email"`
    Password   string `json:"password"`
    AuthorityLevel authorityLevel `json:"authorityLevel"`

}


func (b *User) TableName() string {
    return "user"
   }