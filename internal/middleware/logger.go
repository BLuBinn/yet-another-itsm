package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ZerologMiddleware returns a gin.HandlerFunc for logging requests using zerolog
func ZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := c.ClientIP()

		// Get status code
		statusCode := c.Writer.Status()

		// Build full path with query parameters
		if raw != "" {
			path = path + "?" + raw
		}

		// Create log entry
		logEntry := log.Info()

		// Add log level based on status code
		switch {
		case statusCode >= 500:
			logEntry = log.Error()
		case statusCode >= 400:
			logEntry = log.Warn()
		case statusCode >= 300:
			logEntry = log.Info()
		default:
			logEntry = log.Info()
		}

		// Add fields to log
		logEntry.
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", statusCode).
			Dur("latency", latency).
			Str("ip", clientIP).
			Str("user_agent", c.Request.UserAgent()).
			Int("body_size", c.Writer.Size())

		// Add errors if any
		if len(c.Errors) > 0 {
			logEntry.Strs("errors", c.Errors.Errors())
		}

		logEntry.Msg("HTTP Request")
	}
}

// Recovery middleware with zerolog
func RecoveryWithZerolog() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Error().
			Interface("panic", recovered).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Msg("Panic recovered")

		c.AbortWithStatus(500)
	})
}
