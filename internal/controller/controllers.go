package controller

import (
	"yet-another-itsm/internal/service"
)

type Controllers struct {
	Health       *HealthController
	BusinessUnit *BusinessUnitController
	Department   *DepartmentController
	User         *UserController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		Health:       NewHealthController(services),
		BusinessUnit: NewBusinessUnitController(services),
		Department:   NewDepartmentController(services),
		User:         NewUserController(services),
	}
}
