package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HsiaoCz/stream-media/conf"
	"github.com/HsiaoCz/stream-media/storage"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

type Server struct {
	listenAddr string
	user       *user
	comments   *comments
	store      *storage.Storage
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		listenAddr: fmt.Sprintf("%s:%s", conf.Conf.AppConfig.AppAddr, conf.Conf.AppConfig.AppPort),
		store:      store,
		user:       newUser(),
		comments:   newComments(),
	}
}

func (s *Server) Start() error {
	// user router
	s.user.registerRouter()
	// comments router
	s.comments.registerRouter()
	srv := http.Server{
		Handler:      r,
		Addr:         s.listenAddr,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return srv.ListenAndServe()
}
