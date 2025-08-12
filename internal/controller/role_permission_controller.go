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

type RolePermissionController struct {
	services *service.Services
}

func NewRolePermissionController(services *service.Services) *RolePermissionController {
	return &RolePermissionController{
		services: services,
	}
}

// GetPermissionsByRole godoc
// @Summary Get role permissions
// @Description Get all permissions assigned to a specific role
// @Tags role-permissions
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Success 200 {object} responseModel.RolePermissionsListResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/roles/{roleId}/permissions [get]
func (rpc *RolePermissionController) GetPermissionsByRole(c *gin.Context) {
	log.Info().
		Str("controller", "RolePermissionController").
		Str("endpoint", "GetPermissionsByRole").
		Str("method", c.Request.Method).
		Msg("Get role permissions endpoint called")

	roleID := c.Param("roleId")
	if roleID == "" {
		utils.SendBadRequest(c, constants.ErrRoleIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	rolePermissions, err := rpc.services.RolePermission.GetPermissionsByRole(ctx, roleID)
	if err != nil {
		log.Error().Err(err).Str("role_id", roleID).Msg("Failed to get role permissions")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrieveRolePermissionsMsg)
		return
	}

	log.Info().
		Str("role_id", roleID).
		Int("count", len(rolePermissions)).
		Msg("Successfully retrieved role permissions")

	var rolePermissionResponses []responseModel.RolePermissionDetailResponse
	for _, rolePermission := range rolePermissions {
		rolePermissionResponses = append(rolePermissionResponses, *rolePermission)
	}

	responseData := responseModel.NewRolePermissionsListResponse(
		rolePermissionResponses,
		1,
		len(rolePermissions),
		int64(len(rolePermissions)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved role permissions", responseData)
}

// GetRolePermissionByID godoc
// @Summary Get role permission by ID
// @Description Get a specific role permission by ID
// @Tags role-permissions
// @Accept json
// @Produce json
// @Param rolePermissionId path string true "Role Permission ID"
// @Success 200 {object} responseModel.RolePermissionDetailResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/role-permissions/{rolePermissionId} [get]
func (rpc *RolePermissionController) GetRolePermissionByID(c *gin.Context) {
	log.Info().
		Str("controller", "RolePermissionController").
		Str("endpoint", "GetRolePermissionByID").
		Str("method", c.Request.Method).
		Msg("Get role permission by ID endpoint called")

	id := c.Param("rolePermissionId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrRolePermissionIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	rolePermission, err := rpc.services.RolePermission.GetRolePermissionByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("role_permission_id", id).Msg("Failed to get role permission by ID")
		utils.SendNotFound(c, constants.ErrRolePermissionNotFoundMsg)
		return
	}

	log.Info().
		Str("role_permission_id", id).
		Str("role_name", rolePermission.RoleName).
		Str("permission_name", rolePermission.PermissionName).
		Msg("Successfully retrieved role permission by ID")

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved role permission by ID", rolePermission)
}

// CreateRolePermission godoc
// @Summary Create a new role permission
// @Description Assign a permission to a role
// @Tags role-permissions
// @Accept json
// @Produce json
// @Param rolePermission body responseModel.CreateRolePermissionRequest true "Role Permission data"
// @Success 201 {object} responseModel.RolePermissionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/role-permissions [post]
func (rpc *RolePermissionController) CreateRolePermission(c *gin.Context) {
	var req responseModel.CreateRolePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request body")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "RolePermissionController").
		Str("method", "CreateRolePermission").
		Str("role_id", req.RoleID).
		Str("permission_id", req.PermissionID).
		Msg("Creating new role permission")

	ctx := c.Request.Context()

	rolePermission, err := rpc.services.RolePermission.CreateRolePermission(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create role permission")
		utils.SendInternalServerError(c, constants.ErrFailedToCreateRolePermissionMsg)
		return
	}

	log.Info().
		Str("role_permission_id", rolePermission.ID).
		Str("role_id", rolePermission.RoleID).
		Str("permission_id", rolePermission.PermissionID).
		Msg("Successfully created role permission")

	utils.SendSuccess(c, http.StatusCreated, "Successfully created role permission", rolePermission)
}
