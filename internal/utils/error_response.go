package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yet-another-itsm/internal/constants"
)

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Detail string `json:"detail,omitempty"`
	Code   string `json:"code,omitempty"`
}

// Error creates an error response
func Error(message, code string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Detail: message,
			Code:   code,
		},
	}
}

// ErrorWithContext creates an error response with request context
func ErrorWithContext(c *gin.Context, message, code string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Detail: message,
			Code:   code,
		},
	}
}

// InternalServerError creates a 500 internal server error response
func InternalServerError(message string) *ErrorResponse {
	if message == "" {
		message = "Internal server error"
	}
	return Error(message, constants.ErrCodeInternal)
}

// InternalServerErrorWithContext creates a 500 internal server error response with context
func InternalServerErrorWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Internal server error"
	}
	return ErrorWithContext(c, message, constants.ErrCodeInternal)
}

// NotFound creates a 404 not found error response
func NotFound(message string) *ErrorResponse {
	if message == "" {
		message = "Resource not found"
	}
	return Error(message, constants.ErrCodeNotFound)
}

// NotFoundWithContext creates a 404 not found error response with context
func NotFoundWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Resource not found"
	}
	return ErrorWithContext(c, message, constants.ErrCodeNotFound)
}

// BadRequest creates a 400 bad request error response
func BadRequest(message string) *ErrorResponse {
	if message == "" {
		message = "Bad request"
	}
	return Error(message, constants.ErrCodeBadRequest)
}

// BadRequestWithContext creates a 400 bad request error response with context
func BadRequestWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Bad request"
	}
	return ErrorWithContext(c, message, constants.ErrCodeBadRequest)
}

// Unauthorized creates a 401 unauthorized error response
func Unauthorized(message string) *ErrorResponse {
	if message == "" {
		message = "Unauthorized"
	}
	return Error(message, constants.ErrCodeUnauthorized)
}

// UnauthorizedWithContext creates a 401 unauthorized error response with context
func UnauthorizedWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Unauthorized"
	}
	return ErrorWithContext(c, message, constants.ErrCodeUnauthorized)
}

// Forbidden creates a 403 forbidden error response
func Forbidden(message string) *ErrorResponse {
	if message == "" {
		message = "Forbidden"
	}
	return Error(message, constants.ErrCodeForbidden)
}

// ForbiddenWithContext creates a 403 forbidden error response with context
func ForbiddenWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Forbidden"
	}
	return ErrorWithContext(c, message, constants.ErrCodeForbidden)
}

// Conflict creates a 409 conflict error response
func Conflict(message string) *ErrorResponse {
	if message == "" {
		message = "Conflict"
	}
	return Error(message, constants.ErrCodeConflict)
}

// ConflictWithContext creates a 409 conflict error response with context
func ConflictWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Conflict"
	}
	return ErrorWithContext(c, message, constants.ErrCodeConflict)
}

// ValidationError creates a 422 validation error response
func ValidationError(message string) *ErrorResponse {
	if message == "" {
		message = "Validation failed"
	}
	return Error(message, constants.ErrCodeValidation)
}

// ValidationErrorWithContext creates a 422 validation error response with context
func ValidationErrorWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Validation failed"
	}
	return ErrorWithContext(c, message, constants.ErrCodeValidation)
}

// ServiceUnavailable creates a 503 service unavailable error response
func ServiceUnavailable(message string) *ErrorResponse {
	if message == "" {
		message = "Service unavailable"
	}
	return Error(message, constants.ErrCodeUnavailable)
}

// ServiceUnavailableWithContext creates a 503 service unavailable error response with context
func ServiceUnavailableWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Service unavailable"
	}
	return ErrorWithContext(c, message, constants.ErrCodeUnavailable)
}

// Timeout creates a 408 timeout error response
func Timeout(message string) *ErrorResponse {
	if message == "" {
		message = "Request timeout"
	}
	return Error(message, constants.ErrCodeTimeout)
}

// TimeoutWithContext creates a 408 timeout error response with context
func TimeoutWithContext(c *gin.Context, message string) *ErrorResponse {
	if message == "" {
		message = "Request timeout"
	}
	return ErrorWithContext(c, message, constants.ErrCodeTimeout)
}

// SendInternalServerError sends a 500 internal server error response
func SendInternalServerError(c *gin.Context, message string) {
	SendError(c, http.StatusInternalServerError, message, constants.ErrInternalServerMsg, constants.ErrCodeInternal)
}

// SendNotFound sends a 404 not found error response
func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message, constants.ErrNotFoundMsg, constants.ErrCodeNotFound)
}

// SendBadRequest sends a 400 bad request error response
func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message, constants.ErrBadRequestMsg, constants.ErrCodeBadRequest)
}

// SendUnauthorized sends a 401 unauthorized error response
func SendUnauthorized(c *gin.Context, message string) {
	SendError(c, http.StatusUnauthorized, message, constants.ErrUnauthorizedMsg, constants.ErrCodeUnauthorized)
}

// SendForbidden sends a 403 forbidden error response
func SendForbidden(c *gin.Context, message string) {
	SendError(c, http.StatusForbidden, message, constants.ErrForbiddenMsg, constants.ErrCodeForbidden)
}

// SendConflict sends a 409 conflict error response
func SendConflict(c *gin.Context, message string) {
	SendError(c, http.StatusConflict, message, constants.ErrConflictMsg, constants.ErrCodeConflict)
}

// SendValidationError sends a 422 validation error response
func SendValidationError(c *gin.Context, message string) {
	SendError(c, http.StatusUnprocessableEntity, message, constants.ErrValidationMsg, constants.ErrCodeValidation)
}

// SendServiceUnavailable sends a 503 service unavailable error response
func SendServiceUnavailable(c *gin.Context, message string) {
	SendError(c, http.StatusServiceUnavailable, message, constants.ErrUnavailableMsg, constants.ErrCodeUnavailable)
}

// SendTimeout sends a 408 timeout error response
func SendTimeout(c *gin.Context, message string) {
	SendError(c, http.StatusRequestTimeout, message, constants.ErrTimeoutMsg, constants.ErrCodeTimeout)
}

// SendMethodNotAllowed sends a 405 method not allowed error response
func SendMethodNotAllowed(c *gin.Context, message string) {
	SendError(c, http.StatusMethodNotAllowed, message, constants.ErrMethodNotAllowedMsg, constants.ErrCodeMethodNotAllowed)
}
