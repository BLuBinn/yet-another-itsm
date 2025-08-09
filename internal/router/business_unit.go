package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type BusinessUnitRouter struct {
	controller *controller.BusinessUnitController
	config     *config.Config
}

func NewBusinessUnitRouter(controller *controller.BusinessUnitController, config *config.Config) *BusinessUnitRouter {
	return &BusinessUnitRouter{
		controller: controller,
		config:     config,
	}
}

func (bur *BusinessUnitRouter) SetupBusinessUnitRoutes(v1 *gin.RouterGroup) {
	businessUnitGroup := v1.Group("/business-units").Use(middleware.AuthMiddleWare(&bur.config.OAuth))
	{
		businessUnitGroup.GET("/", bur.controller.GetAllBusinessUnitsInTenant)
		businessUnitGroup.GET("/domain/:domainName", bur.controller.GetBusinessUnitByDomainName)
		businessUnitGroup.GET("/:businessUnitId", bur.controller.GetBusinessUnitByID)
	}
}
