package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(WebApp)

type WebApp struct {
	MySQLConf MysqlConf `mapstructure:"mysql"`
}

type MysqlConf struct {
	Mysql_User string `mapstructure:"mysql_user"`
	Password   string `mapstructure:"password"`
	Mysql_Host string `mapstructure:"mysql_host"`
	Mysql_port string `mapstructure:"mysql_port"`
	DB_Name    string `mapstructure:"db_name"`
}

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("the config is changed", e.Name)
		err = viper.Unmarshal(Conf)
		if err != nil {
			log.Fatal(err)
		}
	})
	return err
}
