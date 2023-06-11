package storage

import (
	"github.com/HsiaoCz/stream-media/conf"
)

type MySqlStorage interface {
	UserLogin() error
	UserRegister() error
	UserDelete() error
	GetUserByName() error
}
type Mysql_Storage struct {
	mysql_user     string
	mysql_password string
	mysql_Host     string
	mysql_port     string
	db_Name        string
}

func NewMysqlStorage() *Mysql_Storage {
	mysql_conf := conf.Conf.MySQLConf
	return &Mysql_Storage{
		mysql_user:     mysql_conf.Mysql_User,
		mysql_password: mysql_conf.Password,
		mysql_Host:     mysql_conf.Mysql_Host,
		mysql_port:     mysql_conf.Mysql_port,
		db_Name:        mysql_conf.DB_Name,
	}
}

func (s *Mysql_Storage) UserRegister() error  { return nil }
func (s *Mysql_Storage) UserLogin() error     { return nil }
func (s *Mysql_Storage) GetUserByName() error { return nil }
func (s *Mysql_Storage) UserDelete() error    { return nil }
