package controller

import (
	"net/http"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DepartmentController struct {
	services *service.Services
}

func NewDepartmentController(services *service.Services) *DepartmentController {
	return &DepartmentController{
		services: services,
	}
}

// GetDepartmentByID godoc
// @Summary Get department by ID
// @Description Get department by ID
// @Tags departments
// @Accept json
// @Produce json
// @Param departmentId path string true "Department ID"
// @Success 200 {object} responseModel.DepartmentResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/departments/{departmentId} [get]
func (dc *DepartmentController) GetDepartmentByID(c *gin.Context) {
	log.Info().
		Str("controller", "DepartmentController").
		Str("endpoint", "GetDepartmentByID").
		Str("method", c.Request.Method).
		Msg("Get department by ID endpoint called")

	id := c.Param("departmentId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrDepartmentIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	department, err := dc.services.Department.GetDepartmentByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToGetDepartmentByIDMsg)
		utils.SendNotFound(c, constants.ErrDepartmentNotFoundMsg)
		return
	}

	log.Info().
		Str("id", id).
		Str("department_name", department.Name).
		Msg(constants.SuccessMsgGetDepartmentByID)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetDepartmentByID, department.ToResponse())
}

// GetDepartmentByName godoc
// @Summary Get department by name
// @Description Get department by name within a business unit
// @Tags departments
// @Accept json
// @Produce json
// @Param business_unit_id query string true "Business Unit ID"
// @Param name query string true "Department Name"
// @Success 200 {object} responseModel.DepartmentResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/departments/ [get]
func (dc *DepartmentController) GetDepartmentByName(c *gin.Context) {
	log.Info().
		Str("controller", "DepartmentController").
		Str("endpoint", "GetDepartmentByName").
		Str("method", c.Request.Method).
		Msg("Get department by name endpoint called")

	businessUnitID := c.Query("business_unit_id")
	name := c.Query("name")

	if businessUnitID == "" {
		utils.SendBadRequest(c, constants.ErrBusinessUnitIDRequiredMsg)
		return
	}

	if name == "" {
		utils.SendBadRequest(c, constants.ErrDepartmentNameRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	department, err := dc.services.Department.GetDepartmentByName(ctx, name)
	if err != nil {
		log.Error().Err(err).Str("name", name).Str("business_unit_id", businessUnitID).Msg(constants.ErrFailedToGetDepartmentByNameMsg)
		utils.SendNotFound(c, constants.ErrDepartmentNotFoundMsg)
		return
	}

	log.Info().
		Str("name", name).
		Str("business_unit_id", businessUnitID).
		Str("department_id", department.ID).
		Msg(constants.SuccessMsgGetDepartmentByName)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetDepartmentByName, department.ToResponse())
}
