package controller

import (
	"yet-another-itsm/internal/service"
)

type Controllers struct {
	Health       *HealthController
	Graph        *GraphController
	BusinessUnit *BusinessUnitController
	Department   *DepartmentController
	User         *UserController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		Health:       NewHealthController(services),
		Graph:        NewGraphController(services),
		BusinessUnit: NewBusinessUnitController(services),
		Department:   NewDepartmentController(services),
		User:         NewUserController(services),
	}
}
