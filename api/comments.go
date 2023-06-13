package api

import (
	"github.com/HsiaoCz/stream-media/storage"
	"github.com/gin-gonic/gin"
)

type comments struct {
	store *storage.Storage
}

func newComments(store *storage.Storage) *comments {
	return &comments{store: store}
}

func (c *comments) registerRouter() {
	r.GET("/videos/:vid-id/comments", c.handleShowComments)
	r.POST("/videos/:vid-id/comments", c.handlePostAComment)
	r.DELETE("/videos/:vid-id/comments/:comments-id", c.handleDeleteAComments)
}

func (c *comments) handleShowComments(ctx *gin.Context)    {}
func (c *comments) handlePostAComment(ctx *gin.Context)    {}
func (c *comments) handleDeleteAComments(ctx *gin.Context) {}
