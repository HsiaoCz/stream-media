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
	user     *user
	comments *comments
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		user:     newUser(store),
		comments: newComments(store),
	}
}

func (s *Server) Start() error {
	// user router
	s.user.registerRouter()
	// comments router
	s.comments.registerRouter()
	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", conf.Conf.AppConfig.AppAddr, conf.Conf.AppConfig.AppPort),
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return srv.ListenAndServe()
}
