package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/goJob/handler/dto/request"
	"github.com/murilogilfelpeto/goJob/handler/dto/response"
	"github.com/murilogilfelpeto/goJob/handler/mapper"
	"github.com/murilogilfelpeto/goJob/schemas"
	"github.com/murilogilfelpeto/goJob/service"
	"net/http"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.OpeningRequestDto{})
)

func CreateOpening(context *gin.Context) {
	var requestBody request.OpeningRequestDto
	err := context.BindJSON(&requestBody)
	if err != nil {
		logger.Errorf("Error while binding JSON: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while binding JSON",
			Timestamp: time.Now(),
			Field:     customValidator.DecryptErrors(err),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
	}

	var opening schemas.Opening = mapper.OpeningRequestToOpening(requestBody)
	opening, err = service.Save(opening)
	if err != nil {
		logger.Errorf("Error while saving opening: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while saving opening: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
	}
	mapper.OpeningToOpeningResponse(opening)
	context.IndentedJSON(http.StatusCreated, opening)
}
