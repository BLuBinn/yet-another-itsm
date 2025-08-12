package service

import (
	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/database"
	"yet-another-itsm/internal/repository"
)

type Services struct {
	Health         HealthService
	Graph          *GraphService
	BusinessUnit   BusinessUnitService
	Department     DepartmentService
	User           UserService
	Role           RoleService
	Permission     PermissionService
	Scope          ScopeService
	RolePermission RolePermissionService
	RoleAssignment RoleAssignmentService
}

func NewServices(db *database.Database, repository *repository.Queries, config *config.Config) *Services {
	return &Services{
		Health:         NewHealthService(db),
		Graph:          NewGraphService(&config.OAuth),
		BusinessUnit:   NewBusinessUnitService(repository),
		Department:     NewDepartmentService(repository),
		User:           NewUserService(repository),
		Role:           NewRoleService(repository),
		Permission:     NewPermissionService(repository),
		Scope:          NewScopeService(repository),
		RolePermission: NewRolePermissionService(repository),
		RoleAssignment: NewRoleAssignmentService(repository),
	}
}
