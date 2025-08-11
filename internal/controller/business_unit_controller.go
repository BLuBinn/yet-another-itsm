package controller

import (
	"net/http"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	responseModel "yet-another-itsm/internal/dtos"
)

type BusinessUnitController struct {
	services *service.Services
}

func NewBusinessUnitController(services *service.Services) *BusinessUnitController {
	return &BusinessUnitController{
		services: services,
	}
}

// GetAllBusinessUnitsInTenant godoc
// @Summary Get all business units in tenant
// @Description Get all business units in tenant
// @Tags business-units
// @Accept json
// @Produce json
// @Param tenantId path string true "Tenant ID"
// @Success 200 {object} responseModel.BusinessUnitsListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/business-units [get]
func (bc *BusinessUnitController) GetAllBusinessUnitsInTenant(c *gin.Context) {
	log.Info().
		Str("controller", "BusinessUnitController").
		Str("endpoint", "GetAllBusinessUnitsInTenant").
		Str("method", c.Request.Method).
		Msg("Get all business units in tenant endpoint called")

	ctx := c.Request.Context()

	businessUnits, err := bc.services.BusinessUnit.GetAllBusinessUnitsInTenant(ctx)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToRetrieveBusinessUnitsMsg)
		utils.SendNotFound(c, constants.ErrBusinessUnitNotFoundMsg)
		return
	}

	log.Info().
		Int("count", len(businessUnits)).
		Msg("Successfully retrieved business units")

	var businessUnitResponses []dtos.BusinessUnitResponse
	for _, bu := range businessUnits {
		businessUnitResponses = append(businessUnitResponses, *bu.ToResponse())
	}

	responseData := responseModel.NewBusinessUnitsListResponse(
		businessUnitResponses,
		1,
		len(businessUnitResponses),
		int64(len(businessUnitResponses)),
	)

	log.Info().
		Int("count", len(businessUnitResponses)).
		Msg(constants.SuccessMsgGetAllBusinessUnits)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetAllBusinessUnits, responseData)
}

// GetBusinessUnitByDomainName godoc
// @Summary Get business unit by domain name
// @Description Get business unit by domain name
// @Tags business-units
// @Accept json
// @Produce json
// @Param domain query string true "Domain Name"
// @Success 200 {object} responseModel.BusinessUnitResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/business-units [get]
func (bc *BusinessUnitController) GetBusinessUnitByDomainName(c *gin.Context) {
	log.Info().
		Str("controller", "BusinessUnitController").
		Str("endpoint", "GetBusinessUnitByDomainName").
		Str("method", c.Request.Method).
		Msg("Get business unit by domain name endpoint called")

	domainName := c.Query("domain")
	if domainName == "" {
		utils.SendBadRequest(c, constants.ErrDomainNameRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	businessUnit, err := bc.services.BusinessUnit.GetBusinessUnitByDomainName(ctx, domainName)
	if err != nil {
		log.Error().Err(err).Str("domain_name", domainName).Msg(constants.ErrFailedToGetBusinessUnitByDomainNameMsg)
		utils.SendNotFound(c, constants.ErrBusinessUnitNotFoundMsg)
		return
	}

	log.Info().
		Str("domain_name", domainName).
		Str("business_unit_id", businessUnit.ID.String()).
		Msg(constants.SuccessMsgGetBusinessUnitByDomainName)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetBusinessUnitByDomainName, businessUnit.ToResponse())
}

// GetBusinessUnitByID godoc
// @Summary Get business unit by ID
// @Description Get business unit by ID
// @Tags business-units
// @Accept json
// @Produce json
// @Param businessUnitId path string true "Business Unit ID"
// @Success 200 {object} responseModel.BusinessUnitResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/business-units/{businessUnitId} [get]
func (bc *BusinessUnitController) GetBusinessUnitByID(c *gin.Context) {
	log.Info().
		Str("controller", "BusinessUnitController").
		Str("endpoint", "GetBusinessUnitByID").
		Str("method", c.Request.Method).
		Msg("Get business unit by ID endpoint called")

	id := c.Param("businessUnitId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrBusinessUnitIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	businessUnit, err := bc.services.BusinessUnit.GetBusinessUnitByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToGetBusinessUnitByIDMsg)
		utils.SendNotFound(c, constants.ErrBusinessUnitNotFoundMsg)
		return
	}

	log.Info().
		Str("id", id).
		Str("business_unit_name", businessUnit.Name).
		Msg(constants.SuccessMsgGetBusinessUnitByID)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetBusinessUnitByID, businessUnit.ToResponse())
}

// GetAllDepartmentsInBusinessUnit godoc
// @Summary Get all departments in business unit
// @Description Get all departments in business unit
// @Tags departments
// @Accept json
// @Produce json
// @Param businessUnitId path string true "Business Unit ID"
// @Success 200 {object} responseModel.DepartmentsListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/business-units/{businessUnitId}/departments [get]
func (bc *BusinessUnitController) GetAllDepartmentsInBusinessUnit(c *gin.Context) {
	log.Info().
		Str("controller", "BusinessUnitController").
		Str("endpoint", "GetAllDepartmentsInBusinessUnit").
		Str("method", c.Request.Method).
		Msg("Get all departments in business unit endpoint called")

	businessUnitID := c.Param("businessUnitId")
	if businessUnitID == "" {
		utils.SendBadRequest(c, constants.ErrBusinessUnitIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	departments, err := bc.services.BusinessUnit.GetAllDepartmentsInBusinessUnit(ctx, businessUnitID)
	if err != nil {
		log.Error().Err(err).Str("business_unit_id", businessUnitID).Msg(constants.ErrFailedToRetrieveDepartmentsMsg)
		utils.SendNotFound(c, constants.ErrDepartmentNotFoundMsg)
		return
	}

	log.Info().
		Int("count", len(departments)).
		Str("business_unit_id", businessUnitID).
		Msg(constants.SuccessMsgGetAllDepartments)

	var departmentResponses []dtos.DepartmentResponse
	for _, dept := range departments {
		departmentResponses = append(departmentResponses, *dept.ToResponse())
	}

	response := responseModel.NewDepartmentsListResponse(
		departmentResponses,
		1,
		len(departmentResponses),
		int64(len(departmentResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetAllDepartments, response)
}
