package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, code, message string, details interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error: &APIError{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string, details interface{}) {
	ErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", message, details)
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", message, nil)
}

// Forbidden sends a 403 Forbidden response
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, "FORBIDDEN", message, nil)
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", message, nil)
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", message, nil)
}

