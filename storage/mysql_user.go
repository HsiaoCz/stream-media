package storage

import "gorm.io/gorm"

type User_Mysql interface {
	UserLogin() error
	UserRegister() error
	UserDelete() error
	GetUserByName() error
}

type user_mysql struct {
	db *gorm.DB
}

func new_user_mysql() *user_mysql {
	return &user_mysql{
		db: &gorm.DB{},
	}
}

func (u *user_mysql) UserLogin() error     { return nil }
func (u *user_mysql) UserRegister() error  { return nil }
func (u *user_mysql) UserDelete() error    { return nil }
func (u *user_mysql) GetUserByName() error { return nil }
