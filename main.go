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

	ms := storage.NewMysqlStorage()
	istore := storage.NewIstore()
	if err := istore.StartConnStorage(ms); err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage()

	srv := api.NewServer(storage)
	log.Println("the server is running on port:", fmt.Sprintf("%s:%s", conf.Conf.AppConfig.AppAddr, conf.Conf.AppConfig.AppPort))
	log.Fatal(srv.Start())
}
