package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/rs/zerolog/log"
)

// RoleAssignmentService defines the interface for role assignment operations
type RoleAssignmentService interface {
	GetUserRoleAssignments(ctx context.Context, userID string) ([]*dtos.RoleAssignment, error)
}

type roleAssignmentService struct {
	repo *repository.Queries
}

func NewRoleAssignmentService(repo *repository.Queries) RoleAssignmentService {
	return &roleAssignmentService{
		repo: repo,
	}
}

func (s *roleAssignmentService) GetUserRoleAssignments(ctx context.Context, userID string) ([]*dtos.RoleAssignment, error) {
	log.Info().
		Str("service", "RoleAssignmentService").
		Str("method", "GetUserRoleAssignments").
		Str("user_id", userID).
		Msg("Getting user role assignments")

	uuid, err := utils.ParseUUID(userID)

	if err != nil {
		log.Error().Err(err).Str("user_id", userID).Msg("Invalid user ID format")
		return nil, fmt.Errorf("%s: %w", constants.ErrInvalidUUIDFormat, err)
	}

	roleAssignments, err := s.repo.GetUserRoleAssignments(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("user_id", userID).Msg("Failed to get role assignments from repository")
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetRoleAssignments, err)
	}

	result := make([]*dtos.RoleAssignment, len(roleAssignments))
	for i, ra := range roleAssignments {
		result[i] = &dtos.RoleAssignment{
			ID:                ra.ID,
			RolePermissionsID: ra.RolePermissionsID.String(),
			AssigneeID:        ra.AssigneeID.String(),
			BusinessUnitID:    ra.BusinessUnitID.String(),
			DepartmentID:      ra.DepartmentID.String(),
			AssignedBy:        ra.AssignedBy.String(),
			AssignedAt:        utils.FormatTime(ra.AssignedAt.Time),
			ExpiresAt:         utils.FormatTime(ra.ExpiresAt.Time),
			Status:            string(ra.Status.StatusEnum),
			RoleName:          ra.RoleName,
			PermissionName:    ra.PermissionName,
			Resource:          ra.Resource,
			Action:            ra.Action,
			UpdatedAt:         utils.FormatTime(ra.UpdatedAt.Time),
		}

		if ra.ScopeName.Valid {
			result[i].ScopeName = ra.ScopeName.String
		}
		if ra.BusinessUnitName.Valid {
			result[i].BusinessUnitName = ra.BusinessUnitName.String
		}
		if ra.DepartmentName.Valid {
			result[i].DepartmentName = ra.DepartmentName.String
		}
		if ra.DeletedAt.Valid {
			result[i].DeletedAt = utils.FormatTime(ra.DeletedAt.Time)
		}
	}

	log.Info().
		Str("service", "RoleAssignmentService").
		Str("method", "GetUserRoleAssignments").
		Str("user_id", userID).
		Int("count", len(result)).
		Msg("Successfully retrieved user role assignments")
	return result, nil
}

func (s *roleAssignmentService) CheckUserPermission(ctx context.Context, userID, resource, action string) (bool, error) {
	log.Info().
		Str("service", "RoleAssignmentService").
		Str("method", "CheckUserPermission").
		Str("user_id", userID).
		Str("resource", resource).
		Str("action", action).
		Msg("Checking user permission")

	uuid, err := utils.ParseUUID(userID)
	if err != nil {
		log.Error().Err(err).Str("user_id", userID).Msg("Invalid user ID format")
		return false, fmt.Errorf("%s: %w", constants.ErrInvalidUUIDFormat, err)
	}

	params := repository.CheckUserPermissionParams{
		AssigneeID: pgtype.UUID{Bytes: uuid, Valid: true},
		Resource:   resource,
		Action:     action,
	}

	hasPermission, err := s.repo.CheckUserPermission(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to check user permission")
		return false, fmt.Errorf("%s: %w", constants.ErrFailedToCheckUserPermission, err)
	}

	log.Info().
		Str("service", "RoleAssignmentService").
		Str("method", "CheckUserPermission").
		Str("user_id", userID).
		Str("resource", resource).
		Str("action", action).
		Bool("has_permission", hasPermission).
		Msg("Successfully checked user permission")
	return hasPermission, nil
}
