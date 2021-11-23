package models

// User Struct

type AuthorityLevel int

const (
	SystemAdmin AuthorityLevel = iota
	EntityAdmin
	SiteAdmin
	SecurityManger
)

type User struct {
	Id             uint           `gorrm:"unique" json:"id"`
	UserName       string         `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	AuthorityLevel AuthorityLevel `json:"authorityLevel"`
}

func (b *User) TableName() string {
	return "user"
}
