package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type GraphRouter struct {
	controller *controller.GraphController
	config     *config.Config
}

func NewGraphRouter(controller *controller.GraphController, config *config.Config) *GraphRouter {
	return &GraphRouter{
		controller: controller,
		config:     config,
	}
}

func (gr *GraphRouter) SetupGraphRoutes(v1 *gin.RouterGroup) {
	graphGroup := v1.Group("/graph").Use(middleware.AuthMiddleWare(&gr.config.OAuth))
	{
		graphGroup.GET("/users/me", gr.controller.GetCurrentUserFromGraph)
	}
}
