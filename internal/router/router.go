package router

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/controller"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Routers struct {
	Health         *HealthRouter
	BusinessUnit   *BusinessUnitRouter
	Department     *DepartmentRouter
	User           *UserRouter
	Role           *RoleRouter
	Permission     *PermissionRouter
	Scope          *ScopeRouter
	RolePermission *RolePermissionRouter
	RoleAssignment *RoleAssignmentRouter
	FormCategory   *FormCategoryRouter
	FormTemplate   *FormTemplateRouter
	FormSection    *FormSectionRouter
}

func NewRouters(controllers *controller.Controllers, config *config.Config) *Routers {
	return &Routers{
		Health:         NewHealthRouter(controllers.Health),
		BusinessUnit:   NewBusinessUnitRouter(controllers.BusinessUnit, config),
		Department:     NewDepartmentRouter(controllers.Department, config),
		User:           NewUserRouter(controllers.User, config),
		Role:           NewRoleRouter(controllers.Role, config),
		Permission:     NewPermissionRouter(controllers.Permission, config),
		Scope:          NewScopeRouter(controllers.Scope, config),
		RolePermission: NewRolePermissionRouter(controllers.RolePermission, config),
		RoleAssignment: NewRoleAssignmentRouter(controllers.RoleAssignment, config),
		FormCategory:   NewFormCategoryRouter(controllers.FormCategory, config),
		FormTemplate:   NewFormTemplateRouter(controllers.FormTemplate, config),
		FormSection:    NewFormSectionRouter(controllers.FormSection, config),
	}
}

func (r *Routers) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	// Health routes
	r.Health.SetupHealthRoutes(v1)

	// Business unit routes
	r.BusinessUnit.SetupBusinessUnitRoutes(v1)

	// Department routes
	r.Department.SetupDepartmentRoutes(v1)

	// User routes
	r.User.SetupUserRoutes(v1)

	// Role routes
	r.Role.SetupRoleRoutes(v1)

	// Permission routes
	r.Permission.SetupPermissionRoutes(v1)

	// Scope routes
	r.Scope.SetupScopeRoutes(v1)

	// Role permission routes
	r.RolePermission.SetupRolePermissionRoutes(v1)

	// Role assignment routes
	r.RoleAssignment.SetupRoleAssignmentRoutes(v1)

	// Form category routes
	r.FormCategory.SetupFormCategoryRoutes(v1)

	// Form template routes
	r.FormTemplate.SetupFormTemplateRoutes(v1)

	// Form section routes
	r.FormSection.SetupFormSectionRoutes(v1)

	router.NoRoute(func(c *gin.Context) {
		log.Warn().
			Str("path", c.Request.URL.Path).
			Str("method", c.Request.Method).
			Str("ip", c.ClientIP()).
			Msg(constants.ErrRouteNotFoundMsg)

		utils.SendNotFound(c, constants.ErrRouteNotFoundMsg)
	})

	router.NoMethod(func(c *gin.Context) {
		log.Warn().
			Str("path", c.Request.URL.Path).
			Str("method", c.Request.Method).
			Str("ip", c.ClientIP()).
			Msg(constants.ErrMethodNotAllowedMsg)

		utils.SendMethodNotAllowed(c, constants.ErrMethodNotAllowedMsg)
	})
}
