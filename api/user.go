package api

import (
	"github.com/HsiaoCz/stream-media/storage"
	"github.com/gin-gonic/gin"
)

type user struct {
	store *storage.Storage
}

func newUser(store *storage.Storage) *user {
	return &user{store: store}
}

func (u *user) registerRouter() {
	r.POST("/user", u.handleUserRegister)
	r.POST("/user/:username", u.handleUserLogin)
	r.GET("/user/:username", u.handleGetUserByName)
	r.DELETE("/user/:username", u.handleUserDelete)
	r.GET("/user/:username/videos", u.handleListAllVideos)
	r.GET("/user/:username/videos/:vid-id", u.handleGetOneVideo)
	r.DELETE("/user/:username/videos/:vid-id", u.handleDeleteOneVideo)
}

func (u *user) handleUserRegister(ctx *gin.Context)   {}
func (u *user) handleUserLogin(ctx *gin.Context)      {}
func (u *user) handleGetUserByName(ctx *gin.Context)  {}
func (u *user) handleUserDelete(ctx *gin.Context)     {}
func (u *user) handleListAllVideos(ctx *gin.Context)  {}
func (u *user) handleGetOneVideo(ctx *gin.Context)    {}
func (u *user) handleDeleteOneVideo(ctx *gin.Context) {}
