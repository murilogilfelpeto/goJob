package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/goJob/handler/dto/request"
	"github.com/murilogilfelpeto/goJob/handler/dto/response"
	"github.com/murilogilfelpeto/goJob/handler/mapper"
	"github.com/murilogilfelpeto/goJob/service"
	"net/http"
	"strconv"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.OpeningRequestDto{})
)

// CreateOpening POST v1/openings
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
		return
	}

	var opening = mapper.OpeningRequestToOpening(requestBody)
	opening, err = service.Save(opening)
	if err != nil {
		logger.Errorf("Error while saving opening: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while saving opening: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}
	responseBody := mapper.OpeningToOpeningResponse(opening)
	context.IndentedJSON(http.StatusCreated, responseBody)
}

func GetAllOpenings(context *gin.Context) {
	openings, err := service.FindAll()
	if err != nil {
		logger.Errorf("Error while getting all openings: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while getting all openings: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
	}

	if len(openings) == 0 {
		logger.Warnf("No openings found")
		context.Status(http.StatusNoContent)
		return
	}

	responseBody := mapper.OpeningsToOpeningsResponse(openings)
	context.IndentedJSON(http.StatusOK, responseBody)
}

func GetOpeningById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logger.Errorf("Error while converting id to int: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while converting id to int: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	opening, err := service.FindById(id)
	if err != nil {
		logger.Errorf("Error while getting opening for id %v: %v", id, err)
		errorResponse := response.ErrorDto{
			Message:   "Error while getting opening for id " + strconv.Itoa(id) + ": " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	responseBody := mapper.OpeningToOpeningResponse(opening)
	context.IndentedJSON(http.StatusOK, responseBody)
}

func UpdateOpening(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logger.Errorf("Error while converting id to int: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while converting id to int: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	requestBody := request.OpeningRequestDto{}
	err = context.BindJSON(&requestBody)
	if err != nil {
		logger.Errorf("Error while binding JSON: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while binding JSON",
			Timestamp: time.Now(),
			Field:     customValidator.DecryptErrors(err),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}
	updatedOpening := mapper.OpeningRequestToOpening(requestBody)
	opening, err := service.UpdateOpening(id, updatedOpening)
	if err != nil {
		errorResponse := response.ErrorDto{
			Message:   "Error while converting id to int: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}
	responseBody := mapper.OpeningToOpeningResponse(opening)
	context.IndentedJSON(http.StatusOK, responseBody)
}

func DeleteOpening(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logger.Errorf("Error while converting id to int: %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while converting id to int: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	err = service.DeleteOpening(id)
	if err != nil {
		logger.Errorf("Error while deleting opening for id %v: %v", id, err)
		errorResponse := response.ErrorDto{
			Message:   "Error while deleting opening for id " + strconv.Itoa(id) + ": " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}
	context.Status(http.StatusOK)
}
