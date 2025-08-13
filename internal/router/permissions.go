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
	const permissionIDPath = "/:permissionId"

	permissionGroup := v1.Group("/permissions").Use(middleware.AuthMiddleWare(&pr.config.OAuth))
	{
		permissionGroup.GET("/", pr.controller.GetAllPermissions)
		permissionGroup.GET(permissionIDPath, pr.controller.GetPermissionByID)
		permissionGroup.GET("/active", pr.controller.GetActivePermissions)
		permissionGroup.GET("/resource/", pr.controller.GetPermissionsByResource)
		permissionGroup.GET("/permission", pr.controller.GetPermissionsByResourceAndAction)
		permissionGroup.POST("/", pr.controller.CreatePermission)
		permissionGroup.PUT(permissionIDPath, pr.controller.UpdatePermission)
		permissionGroup.DELETE(permissionIDPath, pr.controller.DeletePermission)
	}
}
