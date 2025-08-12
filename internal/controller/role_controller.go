package controller

import (
	"net/http"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	responseModel "yet-another-itsm/internal/dtos"
)

type RoleController struct {
	services *service.Services
}

func NewRoleController(services *service.Services) *RoleController {
	return &RoleController{
		services: services,
	}
}

// GetAllRoles godoc
// @Summary Get all roles
// @Description Get all roles in the system
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.RolesListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/roles [get]
func (rc *RoleController) GetAllRoles(c *gin.Context) {
	log.Info().
		Str("controller", "RoleController").
		Str("endpoint", "GetAllRoles").
		Str("method", c.Request.Method).
		Msg("Get all roles endpoint called")

	ctx := c.Request.Context()

	roles, err := rc.services.Role.GetAllRoles(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get roles")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrieveRolesMsg)
		return
	}

	log.Info().
		Int("count", len(roles)).
		Msg("Successfully retrieved roles")

	var roleResponses []responseModel.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, *role.ToResponse())
	}

	responseData := responseModel.NewRolesListResponse(
		roleResponses,
		1,
		len(roleResponses),
		int64(len(roleResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved roles", responseData)
}

// GetRoleByID godoc
// @Summary Get role by ID
// @Description Get a specific role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Success 200 {object} responseModel.RoleResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/roles/{roleId} [get]
func (rc *RoleController) GetRoleByID(c *gin.Context) {
	log.Info().
		Str("controller", "RoleController").
		Str("endpoint", "GetRoleByID").
		Str("method", c.Request.Method).
		Msg("Get role by ID endpoint called")

	id := c.Param("roleId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrRoleIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	role, err := rc.services.Role.GetRoleByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("role_id", id).Msg("Failed to get role by ID")
		utils.SendNotFound(c, constants.ErrRoleNotFoundMsg)
		return
	}

	log.Info().
		Str("role_id", id).
		Str("role_name", role.Name).
		Msg("Successfully retrieved role by ID")

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved role by ID", role.ToResponse())
}

// GetSystemRoles godoc
// @Summary Get system roles
// @Description Get all system roles in the system
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.RolesListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/roles/system [get]
func (rc *RoleController) GetSystemRoles(c *gin.Context) {
	log.Info().
		Str("controller", "RoleController").
		Str("endpoint", "GetSystemRoles").
		Str("method", c.Request.Method).
		Msg("Get system roles endpoint called")

	ctx := c.Request.Context()

	roles, err := rc.services.Role.GetSystemRoles(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get system roles")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrieveRolesMsg)
		return
	}

	log.Info().
		Int("count", len(roles)).
		Msg("Successfully retrieved system roles")

	var roleResponses []responseModel.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, *role.ToResponse())
	}

	responseData := responseModel.NewRolesListResponse(
		roleResponses,
		1,
		len(roleResponses),
		int64(len(roleResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved system roles", responseData)
}

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param role body responseModel.CreateRoleRequest true "Role data"
// @Success 201 {object} responseModel.RoleResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/roles [post]
func (rc *RoleController) CreateRole(c *gin.Context) {
	var req responseModel.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request body")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "RoleController").
		Str("method", "CreateRole").
		Str("role_id", req.ID).
		Str("role_name", req.Name).
		Msg("Creating new role")

	ctx := c.Request.Context()

	role, err := rc.services.Role.CreateRole(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create role")
		utils.SendInternalServerError(c, constants.ErrFailedToCreateRoleMsg)
		return
	}

	log.Info().
		Str("role_id", role.ID).
		Str("role_name", role.Name).
		Msg("Successfully created role")

	utils.SendSuccess(c, http.StatusCreated, "Successfully created role", role.ToResponse())
}
