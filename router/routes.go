package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/goJob/handler"
)

func initRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/v1")
	v1.GET("/ping", ping)
	v1.POST("/openings", handler.CreateOpening)
}

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
