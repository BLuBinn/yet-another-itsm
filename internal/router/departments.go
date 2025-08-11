package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type DepartmentRouter struct {
	controller *controller.DepartmentController
	config     *config.Config
}

func NewDepartmentRouter(controller *controller.DepartmentController, config *config.Config) *DepartmentRouter {
	return &DepartmentRouter{
		controller: controller,
		config:     config,
	}
}

func (dr *DepartmentRouter) SetupDepartmentRoutes(v1 *gin.RouterGroup) {
	departmentGroup := v1.Group("/departments").Use(middleware.AuthMiddleWare(&dr.config.OAuth))
	{
		departmentGroup.GET("/", dr.controller.GetDepartmentByName)
		departmentGroup.GET("/:departmentId", dr.controller.GetDepartmentByID)
	}
}
