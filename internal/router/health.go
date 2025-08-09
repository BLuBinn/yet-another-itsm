package router

import (
	"yet-another-itsm/internal/controller"

	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
	controller *controller.HealthController
}

func NewHealthRouter(controller *controller.HealthController) *HealthRouter {
	return &HealthRouter{
		controller: controller,
	}
}

func (hr *HealthRouter) SetupHealthRoutes(v1 *gin.RouterGroup) {
	healthGroup := v1.Group("/health")
	{
		healthGroup.GET("/", hr.controller.Health)
	}
}
