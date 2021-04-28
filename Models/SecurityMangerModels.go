package Models



type SecurityManger struct {
	
	User `json:"user"`
}


func (b *SecurityManger) TableName() string {
    return "securityManger"
}