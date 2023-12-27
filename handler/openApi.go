package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func OpenApi(c *gin.Context) {
	file, err := os.Open("docs/swagger.json")
	if err != nil {
		logger.Errorf("Error while opening swagger.json: %v", err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func(file *os.File) {
		logger.Infof("Closing swagger.json")
		_ = file.Close()
	}(file)

	c.File("docs/swagger.json")
}
