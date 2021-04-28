package Models
// Our Manger Struct

type authorityLevel int

const ( 
    SystemAdmin authorityLevel = iota
	SiteAdmin
	EntityAdmin
)

type Admin struct {
	
	 
	AuthorityLevel authorityLevel `json:"authorityLevel"`
	User
}



func (b *Admin) TableName() string {
    return "admin"
   }