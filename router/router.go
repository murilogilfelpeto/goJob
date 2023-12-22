package router

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	initRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
