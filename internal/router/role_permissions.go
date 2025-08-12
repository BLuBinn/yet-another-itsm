package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RolePermissionRouter struct {
	controller *controller.RolePermissionController
	config     *config.Config
}

func NewRolePermissionRouter(controller *controller.RolePermissionController, config *config.Config) *RolePermissionRouter {
	return &RolePermissionRouter{
		controller: controller,
		config:     config,
	}
}

func (rpr *RolePermissionRouter) SetupRolePermissionRoutes(v1 *gin.RouterGroup) {
	rolePermissionGroup := v1.Group("/role-permissions").Use(middleware.AuthMiddleWare(&rpr.config.OAuth))
	{
		rolePermissionGroup.GET("/:rolePermissionId", rpr.controller.GetRolePermissionByID)
		rolePermissionGroup.POST("/", rpr.controller.CreateRolePermission)
	}

	roleGroup := v1.Group("/roles").Use(middleware.AuthMiddleWare(&rpr.config.OAuth))
	{
		roleGroup.GET("/:roleId/permissions", rpr.controller.GetPermissionsByRole)
		roleGroup.GET("/:roleId/available-permissions", rpr.controller.GetPermissionsByRole)
	}
}
