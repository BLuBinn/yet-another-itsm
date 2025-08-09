package utils

import "github.com/gin-gonic/gin"

type IAPIResponse interface {
	SetStatus(status string) *APIResponse
	SetMessage(message string) *APIResponse
	SetHeaders(headers interface{}) *APIResponse
	SetData(data interface{}) *APIResponse
	SetStatusCode(statusCode int) *APIResponse
	SetError(error *ErrorResponse) *APIResponse
	Respond(ctx *gin.Context)
}

type APIResponse struct {
	Status     string       `json:"status"`
	Message    string       `json:"message"`
	Headers    interface{}  `json:"headers"`
	Data       interface{}  `json:"data,omitempty"`
	Error      *ErrorDetail `json:"error,omitempty"`
	StatusCode int          `json:"status_code"`
}

func NewAPIResponse() IAPIResponse {
	return &APIResponse{
		Message: "Success",
	}
}

func (r *APIResponse) SetStatus(status string) *APIResponse {
	r.Status = status
	return r
}

func (r *APIResponse) SetMessage(message string) *APIResponse {
	r.Message = message
	return r
}

func (r *APIResponse) SetHeaders(headers interface{}) *APIResponse {
	r.Headers = headers
	return r
}

func (r *APIResponse) SetData(data interface{}) *APIResponse {
	r.Data = data
	return r
}

func (r *APIResponse) SetStatusCode(statusCode int) *APIResponse {
	r.StatusCode = statusCode
	return r
}

func (r *APIResponse) SetError(error *ErrorResponse) *APIResponse {
	r.Error = &error.Error
	return r
}

func (r *APIResponse) Respond(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}

func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	headers := NewHeaders(data, c)

	NewAPIResponse().
		SetStatus("success").
		SetMessage(message).
		SetHeaders(headers).
		SetData(data).
		SetStatusCode(statusCode).
		Respond(c)
}

func SendError(c *gin.Context, statusCode int, detailMessage, message, code string) {
	headers := NewHeaders(nil, c)

	NewAPIResponse().
		SetStatus("error").
		SetMessage(message).
		SetHeaders(headers).
		SetStatusCode(statusCode).
		SetError(Error(detailMessage, code)).
		Respond(c)
}
