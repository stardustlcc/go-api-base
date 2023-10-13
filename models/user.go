package models

type User struct {
	Id       int    `gorm:"primary_key; column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Status   int    `gorm:"default:1;column:status" json:"status"`
}

func (User) TableName() string {
	return "user"
}
