package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/middleware"

	"github.com/gin-gonic/gin"
)

type FormCategoryRouter struct {
	controller *controller.FormCategoryController
	config     *config.Config
}

func NewFormCategoryRouter(controller *controller.FormCategoryController, config *config.Config) *FormCategoryRouter {
	return &FormCategoryRouter{
		controller: controller,
		config:     config,
	}
}

func (fcr *FormCategoryRouter) SetupFormCategoryRoutes(v1 *gin.RouterGroup) {
	formCategoryGroup := v1.Group("/form-categories").Use(middleware.AuthMiddleWare(&fcr.config.OAuth))
	{
		formCategoryGroup.GET("/", fcr.controller.GetFormCategories)
		formCategoryGroup.GET("/:categoryId", fcr.controller.GetFormCategoryByID)
		formCategoryGroup.POST("/", fcr.controller.CreateFormCategory)
		formCategoryGroup.PUT("/:categoryId", fcr.controller.UpdateFormCategory)
		formCategoryGroup.DELETE("/:categoryId", fcr.controller.DeleteFormCategory)
	}
}
