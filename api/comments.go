package api

import (
	"net/http"

	"github.com/HsiaoCz/stream-media/storage"
)

type comments struct {
	store *storage.Storage
}

func newComments(store *storage.Storage) *comments {
	return &comments{store: store}
}

func (c *comments) registerRouter() {
	r.HandleFunc("/videos/{vid-id}/comments", c.handleShowComments).Methods("GET")
	r.HandleFunc("/videos/{vid-id}/comments", c.handlePostAComment).Methods("POST")
	r.HandleFunc("/videos/{vid-id}/comments/{comments-id}", c.handleDeleteAComments).Methods("DELETE")
}

func (c *comments) handleShowComments(w http.ResponseWriter, r *http.Request)    {}
func (c *comments) handlePostAComment(w http.ResponseWriter, r *http.Request)    {}
func (c *comments) handleDeleteAComments(w http.ResponseWriter, r *http.Request) {}
