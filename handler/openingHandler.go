package handler

import "github.com/gin-gonic/gin"

func CreateOpening(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Create Opening",
	})
	return

}
