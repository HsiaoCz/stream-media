package api

import "net/http"

type comments struct{}

func newComments() *comments {
	return &comments{}
}

func (c *comments) registerRouter() {
	r.HandleFunc("/videos/{vid-id}/comments", c.handleShowComments).Methods("GET")
	r.HandleFunc("/videos/{vid-id}/comments", c.handlePostAComment).Methods("POST")
	r.HandleFunc("/videos/{vid-id}/comments/{comments-id}", c.handleDeleteAComments).Methods("DELETE")
}

func (c *comments) handleShowComments(w http.ResponseWriter, r *http.Request)    {}
func (c *comments) handlePostAComment(w http.ResponseWriter, r *http.Request)    {}
func (c *comments) handleDeleteAComments(w http.ResponseWriter, r *http.Request) {}
