package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type FormSectionRouter struct {
	controller *controller.FormSectionController
	config     *config.Config
}

func NewFormSectionRouter(controller *controller.FormSectionController, config *config.Config) *FormSectionRouter {
	return &FormSectionRouter{
		controller: controller,
		config:     config,
	}
}

func (fsr *FormSectionRouter) SetupFormSectionRoutes(v1 *gin.RouterGroup) {
	formSectionGroup := v1.Group("/form-sections").Use(middleware.AuthMiddleWare(&fsr.config.OAuth))
	{
		formSectionGroup.GET("/", fsr.controller.GetFormSections)
		formSectionGroup.GET("/:sectionId", fsr.controller.GetFormSectionByID)
		formSectionGroup.POST("/", fsr.controller.CreateFormSection)
		formSectionGroup.PUT("/:sectionId", fsr.controller.UpdateFormSection)
		formSectionGroup.DELETE("/:sectionId", fsr.controller.DeleteFormSection)
	}
}
