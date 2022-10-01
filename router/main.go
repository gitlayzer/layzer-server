package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

var Router router

type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		GET("/", handler.GetHandler).
		POST("/", handler.PostHandler)
}
