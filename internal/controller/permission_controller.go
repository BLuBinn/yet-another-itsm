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

type PermissionController struct {
	services *service.Services
}

func NewPermissionController(services *service.Services) *PermissionController {
	return &PermissionController{
		services: services,
	}
}

// GetAllPermissions godoc
// @Summary Get all permissions
// @Description Get all permissions in the system
// @Tags permissions
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.PermissionsListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions [get]
func (pc *PermissionController) GetAllPermissions(c *gin.Context) {
	log.Info().
		Str("controller", "PermissionController").
		Str("endpoint", "GetAllPermissions").
		Str("method", c.Request.Method).
		Msg("Get all permissions endpoint called")

	ctx := c.Request.Context()

	permissions, err := pc.services.Permission.GetAllPermissions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get permissions")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}

	log.Info().
		Int("count", len(permissions)).
		Msg("Successfully retrieved permissions")

	var permissionResponses []responseModel.PermissionResponse
	for _, permission := range permissions {
		permissionResponses = append(permissionResponses, *permission.ToResponse())
	}

	responseData := responseModel.NewPermissionsListResponse(
		permissionResponses,
		1,
		len(permissionResponses),
		int64(len(permissionResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved permissions", responseData)
}

// GetPermissionByID godoc
// @Summary Get permission by ID
// @Description Get a specific permission by ID
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} responseModel.PermissionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/{permissionId} [get]
func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	log.Info().
		Str("controller", "PermissionController").
		Str("endpoint", "GetPermissionByID").
		Str("method", c.Request.Method).
		Msg("Get permission by ID endpoint called")

	id := c.Param("permissionId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrPermissionIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	permission, err := pc.services.Permission.GetPermissionByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("permission_id", id).Msg("Failed to get permission by ID")
		utils.SendNotFound(c, constants.ErrPermissionNotFoundMsg)
		return
	}

	log.Info().
		Str("permission_id", id).
		Str("permission_name", permission.Name).
		Msg("Successfully retrieved permission by ID")

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved permission by ID", permission.ToResponse())
}

// GetPermissionsByResource godoc
// @Summary Get permissions by resource
// @Description Get all permissions for a specific resource
// @Tags permissions
// @Accept json
// @Produce json
// @Param resource query string true "Resource"
// @Success 200 {object} responseModel.PermissionsListResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/resource/ [get]
func (pc *PermissionController) GetPermissionsByResource(c *gin.Context) {
	log.Info().
		Str("controller", "PermissionController").
		Str("endpoint", "GetPermissionsByResource").
		Str("method", c.Request.Method).
		Msg("Get permissions by resource endpoint called")

	resource := c.Query("resource")
	if resource == "" {
		utils.SendBadRequest(c, constants.ErrResourceRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	permissions, err := pc.services.Permission.GetPermissionsByResource(ctx, resource)

	exists, err := pc.services.Permission.ResourceExists(ctx, resource)
	if err != nil {
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}
	if !exists {
		utils.SendNotFound(c, constants.ErrResourceNotFoundMsg)
		return
	}

	if len(permissions) == 0 {
		utils.SendNotFound(c, constants.ErrNoPermissionsFoundMsg)
		return
	}

	log.Info().
		Str("resource", resource).
		Int("count", len(permissions)).
		Msg("Successfully retrieved permissions by resource")

	var permissionResponses []responseModel.PermissionResponse
	for _, permission := range permissions {
		permissionResponses = append(permissionResponses, *permission.ToResponse())
	}

	responseData := responseModel.NewPermissionsListResponse(
		permissionResponses,
		1,
		len(permissionResponses),
		int64(len(permissionResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved permissions by resource", responseData)
}

// GetPermissionsByResourceAndAction godoc
// @Summary Get permission by resource and action
// @Description Get a specific permission by resource and action
// @Tags permissions
// @Accept json
// @Produce json
// @Param resource path string true "Resource"
// @Param action path string true "Action"
// @Success 200 {object} responseModel.PermissionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/permission [get]
func (pc *PermissionController) GetPermissionsByResourceAndAction(c *gin.Context) {
	log.Info().
		Str("controller", "PermissionController").
		Str("endpoint", "GetPermissionsByResourceAndAction").
		Str("method", c.Request.Method).
		Msg("Get permission by resource and action endpoint called")

	resource := c.Query("resource")
	action := c.Query("action")

	if resource == "" {
		utils.SendBadRequest(c, constants.ErrResourceRequiredMsg)
		return
	}

	if action == "" {
		utils.SendBadRequest(c, constants.ErrActionRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	resourceExists, err := pc.services.Permission.ResourceExists(ctx, resource)
	if err != nil {
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}
	if !resourceExists {
		utils.SendNotFound(c, constants.ErrResourceNotFoundMsg)
		return
	}

	actionExists, err := pc.services.Permission.ActionExists(ctx, action)
	if err != nil {
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}
	if !actionExists {
		utils.SendNotFound(c, constants.ErrActionNotFoundMsg)
		return
	}

	resourceActionExists, err := pc.services.Permission.ResourceAndActionExists(ctx, resource, action)
	if err != nil {
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}
	if !resourceActionExists {
		utils.SendNotFound(c, constants.ErrResourceActionNotFoundMsg)
		return
	}

	permission, err := pc.services.Permission.GetPermissionsByResourceAndAction(ctx, resource, action)
	if err != nil {
		log.Error().
			Err(err).
			Str("resource", resource).
			Str("action", action).
			Msg("Failed to get permission by resource and action")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}

	log.Info().
		Str("resource", resource).
		Str("action", action).
		Str("permission_name", permission.Name).
		Msg("Successfully retrieved permission by resource and action")

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved permission by resource and action", permission.ToResponse())
}

// GetActivePermissions godoc
// @Summary Get active permissions
// @Description Get all active permissions in the system
// @Tags permissions
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.PermissionsListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/active [get]
func (pc *PermissionController) GetActivePermissions(c *gin.Context) {
	log.Info().
		Str("controller", "PermissionController").
		Str("endpoint", "GetActivePermissions").
		Str("method", c.Request.Method).
		Msg("Get active permissions endpoint called")

	ctx := c.Request.Context()

	permissions, err := pc.services.Permission.GetActivePermissions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get active permissions")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrievePermissionsMsg)
		return
	}

	log.Info().
		Int("count", len(permissions)).
		Msg("Successfully retrieved active permissions")

	var permissionResponses []responseModel.PermissionResponse
	for _, permission := range permissions {
		permissionResponses = append(permissionResponses, *permission.ToResponse())
	}

	responseData := responseModel.NewPermissionsListResponse(
		permissionResponses,
		1,
		len(permissionResponses),
		int64(len(permissionResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved active permissions", responseData)
}

// CreatePermission godoc
// @Summary Create a new permission
// @Description Create a new permission
// @Tags permissions
// @Accept json
// @Produce json
// @Param permission body responseModel.CreatePermissionRequest true "Permission data"
// @Success 201 {object} responseModel.PermissionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions [post]
func (pc *PermissionController) CreatePermission(c *gin.Context) {
	var req responseModel.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request body")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "PermissionController").
		Str("method", "CreatePermission").
		Str("permission_id", req.ID).
		Str("permission_name", req.Name).
		Msg("Creating new permission")

	ctx := c.Request.Context()

	permission, err := pc.services.Permission.CreatePermission(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create permission")
		utils.SendInternalServerError(c, constants.ErrFailedToCreatePermissionMsg)
		return
	}

	log.Info().
		Str("permission_id", permission.ID).
		Str("permission_name", permission.Name).
		Msg("Successfully created permission")

	utils.SendSuccess(c, http.StatusCreated, "Successfully created permission", permission.ToResponse())
}

// UpdatePermission godoc
// @Summary Update a permission
// @Description Update an existing permission
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Param permission body responseModel.UpdatePermissionRequest true "Permission data"
// @Success 200 {object} responseModel.PermissionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/{permissionId} [put]
func (pc *PermissionController) UpdatePermission(c *gin.Context) {
	id := c.Param("permissionId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrPermissionIDRequiredMsg)
		return
	}

	var req responseModel.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request body")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "PermissionController").
		Str("method", "UpdatePermission").
		Str("permission_id", id).
		Str("permission_name", req.Name).
		Msg("Updating permission")

	ctx := c.Request.Context()

	permission, err := pc.services.Permission.UpdatePermission(ctx, id, &req)
	if err != nil {
		log.Error().Err(err).Str("permission_id", id).Msg("Failed to update permission")
		utils.SendInternalServerError(c, constants.ErrFailedToUpdatePermissionMsg)
		return
	}

	log.Info().
		Str("permission_id", permission.ID).
		Str("permission_name", permission.Name).
		Msg("Successfully updated permission")

	utils.SendSuccess(c, http.StatusOK, "Successfully updated permission", permission.ToResponse())
}

// DeletePermission godoc
// @Summary Delete a permission
// @Description Delete an existing permission
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} responseModel.SuccessResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/permissions/{permissionId} [delete]
func (pc *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("permissionId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrPermissionIDRequiredMsg)
		return
	}

	log.Info().
		Str("controller", "PermissionController").
		Str("method", "DeletePermission").
		Str("permission_id", id).
		Msg("Deleting permission")

	ctx := c.Request.Context()

	err := pc.services.Permission.DeletePermission(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("permission_id", id).Msg("Failed to delete permission")
		utils.SendInternalServerError(c, constants.ErrFailedToDeletePermissionMsg)
		return
	}

	log.Info().
		Str("permission_id", id).
		Msg("Successfully deleted permission")

	utils.SendSuccess(c, http.StatusOK, "Successfully deleted permission", nil)
}
