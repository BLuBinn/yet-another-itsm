package utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	ContentType   string `json:"content-type,omitempty"   default:"application/json"`
	ContentLength int    `json:"content-length,omitempty"`
	Method        string `json:"method,omitempty"`
	Date          string `json:"date,omitempty"`
	RequestID     string `json:"request-id,omitempty"`
	Server        string `json:"server,omitempty"`
	Endpoint      string `json:"endpoint,omitempty"`
}

func NewHeaders(data interface{}, ctx *gin.Context) *Headers {
	jsonData, err := json.Marshal(data)
	if err != nil {
		jsonData = []byte{}
	}
	contentLength := len(jsonData)

	date := time.Now().Format(time.RFC1123)
	server := strings.Split(ctx.Request.Host, ":")[0]
	endpoint := ctx.Request.URL.String()
	method := ctx.Request.Method
	requestID := ctx.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	return &Headers{
		Method:        method,
		ContentLength: contentLength,
		RequestID:     requestID,
		Date:          date,
		Server:        server,
		Endpoint:      endpoint,
	}
}
