package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controller.UserController
	config     *config.Config
}

func NewUserRouter(controller *controller.UserController, config *config.Config) *UserRouter {
	return &UserRouter{
		controller: controller,
		config:     config,
	}
}

func (ur *UserRouter) SetupUserRoutes(v1 *gin.RouterGroup) {
	userGroup := v1.Group("/users").Use(middleware.AuthMiddleWare(&ur.config.OAuth))
	{
		userGroup.GET("/:userId", ur.controller.GetUserByID)
		userGroup.GET("/email/:email", ur.controller.GetUserByEmail)
		userGroup.POST("/", ur.controller.CreateUser)
	}

	departmentGroup := v1.Group("/departments").Use(middleware.AuthMiddleWare(&ur.config.OAuth))
	{
		departmentGroup.GET("/:departmentId/users", ur.controller.GetAllUsersInDepartment)
	}
}
