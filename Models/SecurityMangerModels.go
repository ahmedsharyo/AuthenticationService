package Models



type SecurityManger struct {
	gorm.Model
	
	User
}


func (b *SecurityManger) TableName() string {
    return "securityManger"
}