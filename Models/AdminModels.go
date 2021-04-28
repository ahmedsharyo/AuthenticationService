package Models
// Our Manger Struct

type authorityLevel int

const ( 
    SystemAdmin authorityLevel = iota
	SiteAdmin
	EntityAdmin
)

type Admin struct {
	gorm.Model
	
	user User `json:"user"`
	AuthorityLevel authorityLevel `json:"authorityLevel"`
}


func (b *Admin) TableName() string {
    return "admin"
   }