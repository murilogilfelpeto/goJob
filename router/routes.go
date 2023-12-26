package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/goJob/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/ping", ping)
	v1.POST("/opening", handler.CreateOpening)
}

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
