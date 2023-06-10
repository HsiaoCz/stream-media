package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HsiaoCz/stream-media/conf"
	"github.com/HsiaoCz/stream-media/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      *storage.Storage
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		listenAddr: fmt.Sprintf("%s:%s", conf.Conf.AppConfig.AppAddr, conf.Conf.AppConfig.AppPort),
		store:      store,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/user", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/user/{username}", s.handleUserLogin).Methods("POST")
	r.HandleFunc("/user/{username}", s.handleGetUserByName).Methods("GET")
	r.HandleFunc("/user/{username}", s.handleUserDelete).Methods("DETELE")
	srv := http.Server{
		Handler:      r,
		Addr:         s.listenAddr,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return srv.ListenAndServe()
}
