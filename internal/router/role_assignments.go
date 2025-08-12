package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RoleAssignmentRouter struct {
	controller *controller.RoleAssignmentController
	config     *config.Config
}

func NewRoleAssignmentRouter(controller *controller.RoleAssignmentController, config *config.Config) *RoleAssignmentRouter {
	return &RoleAssignmentRouter{
		controller: controller,
		config:     config,
	}
}

func (rar *RoleAssignmentRouter) SetupRoleAssignmentRoutes(v1 *gin.RouterGroup) {
	userGroup := v1.Group("/users").Use(middleware.AuthMiddleWare(&rar.config.OAuth))
	{
		userGroup.GET("/:userId/role-assignments", rar.controller.GetUserRoleAssignments)
	}
}
