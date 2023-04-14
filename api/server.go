package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
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
