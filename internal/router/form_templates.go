package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type FormTemplateRouter struct {
	controller *controller.FormTemplateController
	config     *config.Config
}

func NewFormTemplateRouter(controller *controller.FormTemplateController, config *config.Config) *FormTemplateRouter {
	return &FormTemplateRouter{
		controller: controller,
		config:     config,
	}
}

func (ftr *FormTemplateRouter) SetupFormTemplateRoutes(v1 *gin.RouterGroup) {
	formTemplateGroup := v1.Group("/form-templates").Use(middleware.AuthMiddleWare(&ftr.config.OAuth))
	{
		formTemplateGroup.GET("/", ftr.controller.GetFormTemplates)
		formTemplateGroup.GET("/:templateId", ftr.controller.GetFormTemplateByID)
		formTemplateGroup.GET("/category/:categoryId", ftr.controller.GetFormTemplatesByCategory)
		formTemplateGroup.POST("/", ftr.controller.CreateFormTemplate)
		formTemplateGroup.PUT("/:templateId", ftr.controller.UpdateFormTemplate)
		// formTemplateGroup.POST("/:templateId/publish", ftr.controller.PublishFormTemplate)
		formTemplateGroup.DELETE("/:templateId", ftr.controller.DeleteFormTemplate)
	}
}
