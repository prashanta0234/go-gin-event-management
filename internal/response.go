package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status"`
	Success bool        `json:"success"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
		Status:  http.StatusOK,
		Success: true,
	}
}

func ErrorResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
		Status:  http.StatusBadRequest,
		Success: false,
	}
}

func NotFoundResponse(message string) Response {
	return Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusNotFound,
		Success: false,
	}
}

func UnauthorizedResponse(message string) Response {
	return Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusUnauthorized,
		Success: false,
	}
}

func ServerErrorResponse(message string) Response {
	return Response{
		Message: message,
		Data:    nil,
		Status:  http.StatusInternalServerError,
		Success: false,
	}
}

func ValidationErrorResponse(message string, errors interface{}) Response {
	return Response{
		Message: message,
		Data:    errors,
		Status:  http.StatusUnprocessableEntity,
		Success: false,
	}
}

func SendSuccess(c *gin.Context, message string, data interface{}) {
	response := SuccessResponse(message, data)
	c.JSON(response.Status, response)
}

func SendError(c *gin.Context, message string, data interface{}) {
	response := ErrorResponse(message, data)
	c.JSON(response.Status, response)
}

func SendNotFound(c *gin.Context, message string) {
	response := NotFoundResponse(message)
	c.JSON(response.Status, response)
}

func SendUnauthorized(c *gin.Context, message string) {
	response := UnauthorizedResponse(message)
	c.JSON(response.Status, response)
}

func SendServerError(c *gin.Context, message string) {
	response := ServerErrorResponse(message)
	c.JSON(response.Status, response)
}

func SendValidationError(c *gin.Context, message string, errors interface{}) {
	response := ValidationErrorResponse(message, errors)
	c.JSON(response.Status, response)
}
