package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ScopeRouter struct {
	controller *controller.ScopeController
	config     *config.Config
}

func NewScopeRouter(controller *controller.ScopeController, config *config.Config) *ScopeRouter {
	return &ScopeRouter{
		controller: controller,
		config:     config,
	}
}

func (sr *ScopeRouter) SetupScopeRoutes(v1 *gin.RouterGroup) {
	scopeGroup := v1.Group("/scopes").Use(middleware.AuthMiddleWare(&sr.config.OAuth))
	{
		scopeGroup.GET("/", sr.controller.GetAllScopes)
		scopeGroup.GET("/:scopeId", sr.controller.GetScopeByID)
		scopeGroup.POST("/", sr.controller.CreateScope)
	}
}
