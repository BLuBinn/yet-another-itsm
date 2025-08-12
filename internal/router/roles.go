package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	controller *controller.RoleController
	config     *config.Config
}

func NewRoleRouter(controller *controller.RoleController, config *config.Config) *RoleRouter {
	return &RoleRouter{
		controller: controller,
		config:     config,
	}
}

func (rr *RoleRouter) SetupRoleRoutes(v1 *gin.RouterGroup) {
	roleGroup := v1.Group("/roles").Use(middleware.AuthMiddleWare(&rr.config.OAuth))
	{
		roleGroup.GET("/", rr.controller.GetAllRoles)
		roleGroup.GET("/:roleId", rr.controller.GetRoleByID)
		roleGroup.GET("/system", rr.controller.GetSystemRoles)
		roleGroup.POST("/", rr.controller.CreateRole)
	}
}
