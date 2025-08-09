package controller

import (
	"net/http"
	"time"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	responseModel "yet-another-itsm/internal/dtos"
)

type HealthController struct {
	services *service.Services
}

func NewHealthController(services *service.Services) *HealthController {
	return &HealthController{
		services: services,
	}
}

// Health godoc
// @Summary Health check
// @Description Check the health of the system
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.HealthResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/health [get]
func (h *HealthController) Health(c *gin.Context) {
	log.Info().
		Str("controller", "HealthController").
		Str("endpoint", "health").
		Str("method", c.Request.Method).
		Msg("Health check endpoint called")

	ctx := c.Request.Context()
	health, err := h.services.Health.CheckHealth(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Health check failed")
		utils.SendInternalServerError(c, constants.ErrHealthCheckFailedMsg)
		return
	}

	var statusCode int
	var message string
	switch health.Status {
	case constants.HealthStatusOK:
		statusCode = http.StatusOK
		message = constants.ErrSystemHealthyMsg
	case constants.HealthStatusDegraded:
		statusCode = http.StatusPartialContent
		message = constants.ErrSystemPartiallyHealthyMsg
	default:
		statusCode = http.StatusServiceUnavailable
		message = constants.ErrSystemUnhealthyMsg
	}

	log.Info().
		Str("health_status", health.Status).
		Interface("services", health.Services).
		Msg("Health check completed")

	// Create response data
	data := responseModel.HealthResponse{
		Message:   message,
		Status:    health.Status,
		Services:  health.Services,
		Timestamp: time.Now().UTC(),
	}

	utils.SendSuccess(c, statusCode, constants.SuccessMsgHealthCheck, data)
}
