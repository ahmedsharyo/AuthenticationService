package Models



type SecurityManger struct {
	
	User 
}


func (b *SecurityManger) TableName() string {
    return "securityManger"
}