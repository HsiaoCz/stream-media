package api

import (
	"net/http"

	"github.com/HsiaoCz/stream-media/storage"
)

type user struct {
	store *storage.Storage
}

func newUser(store *storage.Storage) *user {
	return &user{store: store}
}

func (u *user) registerRouter() {
	r.HandleFunc("/user", u.handleUserRegister).Methods("POST")
	r.HandleFunc("/user/{username}", u.handleUserLogin).Methods("POST")
	r.HandleFunc("/user/{username}", u.handleGetUserByName).Methods("GET")
	r.HandleFunc("/user/{username}", u.handleUserDelete).Methods("DETELE")
	r.HandleFunc("/user/{username}/videos", u.handleListAllVideos).Methods("GET")
	r.HandleFunc("/user/{username}/videos/{vid-id}", u.handleGetOneVideo).Methods("GET")
	r.HandleFunc("/user/{username}/videos/{vid-id}", u.handleDeleteOneVideo).Methods("DELETE")
}

func (u *user) handleUserRegister(w http.ResponseWriter, r *http.Request)   {}
func (u *user) handleUserLogin(w http.ResponseWriter, r *http.Request)      {}
func (u *user) handleGetUserByName(w http.ResponseWriter, r *http.Request)  {}
func (u *user) handleUserDelete(w http.ResponseWriter, r *http.Request)     {}
func (u *user) handleListAllVideos(w http.ResponseWriter, r *http.Request)  {}
func (u *user) handleGetOneVideo(w http.ResponseWriter, r *http.Request)    {}
func (u *user) handleDeleteOneVideo(w http.ResponseWriter, r *http.Request) {}
