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

// CreateOpening handles the POST request to /v1/openings.
// It binds the JSON request body to an OpeningRequestDto,
// saves the opening, and returns the created OpeningResponseDto.
// @Summary Create an opening
// @Description Create a new opening
// @Tags openings
// @Accept json
// @Produce json
// @Param requestBody body request.OpeningRequestDto true "Opening request body"
// @Success 201 {object} response.OpeningResponseDto
// @Failure 422 {object} response.ErrorDto
// @Failure 500 {object} response.ErrorDto
// @Router /v1/openings [post]
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

// GetAllOpenings
// @Summary Get all openings
// @Description Retrieve all openings and return them in JSON format
// @Produce json
// @Success 200 {object} response.OpeningResponseDto
// @Failure 204 "No openings found"
// @Failure 500 "Internal server error"
// @Router /openings [get]
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

// GetOpeningById retrieves an opening by its ID
// @Summary Get an opening by ID
// @Description Retrieves an opening by its ID
// @Tags Opening
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} response.OpeningResponseDto
// @Failure 400 {object} response.ErrorDto
// @Failure 500 {object} response.ErrorDto
// @Router /opening/{id} [get]
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

// UpdateOpening is a handler function for updating an opening.
// It expects the ID of the opening in the URL path and the updated opening data in the request body.
// @Summary Update an opening
// @Description Update an existing opening by ID
// @ID update-opening
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Param requestBody body request.OpeningRequestDto true "Updated opening data"
// @Success 200 {object} response.OpeningResponseDto "Updated opening"
// @Failure 400 {object} response.ErrorDto "Bad Request"
// @Failure 422 {object} response.ErrorDto "Unprocessable Entity"
// @Failure 500 {object} response.ErrorDto "Internal Server Error"
// @Router /opening/{id} [put]
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

// @Summary Delete Opening
// @Description Deletes an opening based on the provided ID.
// @Tags Opening
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200
// @Failure 400 {object} response.ErrorDto "Bad Request"
// @Failure 500 {object} response.ErrorDto "Internal Server Error"
// @Router /opening/{id} [delete]
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
