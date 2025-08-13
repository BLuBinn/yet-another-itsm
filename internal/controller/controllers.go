package controller

import (
	"yet-another-itsm/internal/service"
)

type Controllers struct {
	Health         *HealthController
	BusinessUnit   *BusinessUnitController
	Department     *DepartmentController
	User           *UserController
	Role           *RoleController
	Permission     *PermissionController
	Scope          *ScopeController
	RolePermission *RolePermissionController
	RoleAssignment *RoleAssignmentController
	FormCategory   *FormCategoryController
	FormTemplate   *FormTemplateController
	FormSection    *FormSectionController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		Health:         NewHealthController(services),
		BusinessUnit:   NewBusinessUnitController(services),
		Department:     NewDepartmentController(services),
		User:           NewUserController(services),
		Role:           NewRoleController(services),
		Permission:     NewPermissionController(services),
		Scope:          NewScopeController(services),
		RolePermission: NewRolePermissionController(services),
		RoleAssignment: NewRoleAssignmentController(services),
		FormCategory:   NewFormCategoryController(services),
		FormTemplate:   NewFormTemplateController(services),
		FormSection:    NewFormSectionController(services),
	}
}
