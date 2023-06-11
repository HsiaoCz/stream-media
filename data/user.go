package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(40);" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
}

func (u User) TableName() string {
	return "user"
}
