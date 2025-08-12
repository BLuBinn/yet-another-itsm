package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/rs/zerolog/log"
)

type RoleAssignmentController struct {
	services *service.Services
}

func NewRoleAssignmentController(services *service.Services) *RoleAssignmentController {
	return &RoleAssignmentController{
		services: services,
	}
}

// GetUserRoleAssignments godoc
// @Summary Get user role assignments
// @Description Retrieve all role assignments for a specific user
// @Tags role-assignments
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dtos.RoleAssignmentsListResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /role-assignments/user/{userId} [get]
// @Security BearerAuth
func (c *RoleAssignmentController) GetUserRoleAssignments(ctx *gin.Context) {
	userID := ctx.Param("userId")
	if userID == "" {
		log.Ctx(ctx).Error().Msg(constants.ErrUserIDRequiredMsg)
		utils.SendBadRequest(ctx, constants.ErrUserIDRequiredMsg)
		return
	}

	log.Ctx(ctx).Info().Str("user_id", userID).Msg("Getting user role assignments")

	roleAssignments, err := c.services.RoleAssignment.GetUserRoleAssignments(ctx.Request.Context(), userID)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg(constants.ErrFailedToRetrieveRoleAssignmentsMsg)
		utils.SendNotFound(ctx, constants.ErrFailedToRetrieveRoleAssignmentsMsg)
		return
	}

	var detailResponses []dtos.RoleAssignmentDetailResponse
	for _, assignment := range roleAssignments {
		detailResponses = append(detailResponses, dtos.RoleAssignmentDetailResponse{
			ID:                assignment.ID.String(),
			RolePermissionsID: assignment.RolePermissionsID,
			AssigneeID:        assignment.AssigneeID,
			AssigneeName:      assignment.AssigneeName,
			AssigneeEmail:     assignment.AssigneeEmail,
			BusinessUnitID:    assignment.BusinessUnitID,
			BusinessUnitName:  assignment.BusinessUnitName,
			DepartmentID:      assignment.DepartmentID,
			DepartmentName:    assignment.DepartmentName,
			AssignedBy:        assignment.AssignedBy,
			AssignedAt:        assignment.AssignedAt,
			ExpiresAt:         assignment.ExpiresAt,
			RoleName:          assignment.RoleName,
			PermissionName:    assignment.PermissionName,
			Resource:          assignment.Resource,
			Action:            assignment.Action,
			ScopeName:         assignment.ScopeName,
			Status:            assignment.Status,
			UpdatedAt:         assignment.UpdatedAt,
			DeletedAt:         assignment.DeletedAt,
		})
	}

	response := dtos.NewRoleAssignmentsListResponse(detailResponses, 1, len(detailResponses), int64(len(detailResponses)))
	utils.SendSuccess(ctx, http.StatusOK, constants.SuccessMsg, response)
}
