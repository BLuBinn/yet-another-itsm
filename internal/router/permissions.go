package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct {
	controller *controller.PermissionController
	config     *config.Config
}

func NewPermissionRouter(controller *controller.PermissionController, config *config.Config) *PermissionRouter {
	return &PermissionRouter{
		controller: controller,
		config:     config,
	}
}

func (pr *PermissionRouter) SetupPermissionRoutes(v1 *gin.RouterGroup) {
	permissionGroup := v1.Group("/permissions").Use(middleware.AuthMiddleWare(&pr.config.OAuth))
	{
		permissionGroup.GET("/", pr.controller.GetAllPermissions)
		permissionGroup.GET("/:permissionId", pr.controller.GetPermissionByID)
		permissionGroup.GET("/active", pr.controller.GetActivePermissions)
		permissionGroup.GET("/resource/", pr.controller.GetPermissionsByResource)
		permissionGroup.GET("/permission", pr.controller.GetPermissionsByResourceAndAction)
		permissionGroup.POST("/", pr.controller.CreatePermission)
		permissionGroup.PUT("/:permissionId", pr.controller.UpdatePermission)
		permissionGroup.DELETE("/:permissionId", pr.controller.DeletePermission)
	}
}
