package Models
// User Struct
type User struct {
    Id      uint   `json:"id"`
    UserName    string `json:"username"`
    Email   string `json:"email"`
    Password   string `json:"password"`
    
}

func (b *User) TableName() string {
    return "user"
   }