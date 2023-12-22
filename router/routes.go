package router

import "github.com/gin-gonic/gin"

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/ping", ping)
}

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
