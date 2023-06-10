package main

import (
	"fmt"
	"log"

	"github.com/HsiaoCz/stream-media/api"
	"github.com/HsiaoCz/stream-media/conf"
	"github.com/HsiaoCz/stream-media/storage"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	mysql_stroage := storage.NewMysqlStorage()
	redis_storage := storage.NewRedisStorage()

	storage := storage.NewStorage(mysql_stroage, redis_storage)

	srv := api.NewServer(storage)
	log.Println("the server is running on port:", fmt.Sprintf("%s:%s", conf.Conf.AppConfig.AppAddr, conf.Conf.AppConfig.AppPort))
	log.Fatal(srv.Start())
}
