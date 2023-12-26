package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/goJob/handler/dto/request"
	"github.com/murilogilfelpeto/goJob/handler/dto/response"
	"net/http"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.OpeningRequestDto{})
)

func CreateOpening(context *gin.Context) {
	var newOpening request.OpeningRequestDto
	err := context.BindJSON(&newOpening)
	if err != nil {
		logger.Errorf("Error while binding JSON: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while binding JSON",
			Timestamp: time.Now(),
			Field:     customValidator.DecryptErrors(err),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
	}

}
