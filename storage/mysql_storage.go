package storage

import (
	"fmt"

	"github.com/HsiaoCz/stream-media/conf"
	"github.com/HsiaoCz/stream-media/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql_Storage struct {
	mc *mysqlConf
	Um user_Mysql
}

type mysqlConf struct {
	mysql_user     string
	mysql_password string
	mysql_Host     string
	mysql_port     string
	db_Name        string
}

func NewMysqlStorage() *Mysql_Storage {
	myc := conf.Conf.MySQLConf
	return &Mysql_Storage{
		mc: &mysqlConf{
			mysql_user:     myc.Mysql_User,
			mysql_password: myc.Password,
			mysql_Host:     myc.Mysql_Host,
			mysql_port:     myc.Mysql_port,
			db_Name:        myc.DB_Name,
		},
		Um: new_user_mysql(),
	}
}

func (m *Mysql_Storage) initStore() error {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.mc.mysql_user, m.mc.mysql_password, m.mc.mysql_Host, m.mc.mysql_port, m.mc.db_Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&data.User{})
	return err
}
