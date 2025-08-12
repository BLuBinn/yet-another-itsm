package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type RolePermissionService interface {
	GetPermissionsByRole(ctx context.Context, roleID string) ([]*dtos.RolePermissionDetailResponse, error)
	GetRolePermissionByID(ctx context.Context, id string) (*dtos.RolePermissionDetailResponse, error)
	CreateRolePermission(ctx context.Context, req *dtos.CreateRolePermissionRequest) (*dtos.RolePermissionResponse, error)
}

type rolePermissionService struct {
	repo *repository.Queries
}

func NewRolePermissionService(repo *repository.Queries) RolePermissionService {
	return &rolePermissionService{
		repo: repo,
	}
}

func (s *rolePermissionService) GetPermissionsByRole(ctx context.Context, roleID string) ([]*dtos.RolePermissionDetailResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RolePermissionService").
		Str("method", "GetRolePermissions").
		Str("role_id", roleID).
		Str("user_id", userID).
		Msg("Getting role permissions")

	rolePermissions, err := s.repo.GetPermissionsByRole(ctx, roleID)
	if err != nil {
		log.Error().Err(err).Str("role_id", roleID).Msg("Failed to get role permissions from repository")
		return nil, fmt.Errorf("failed to get role permissions from repository: %w", err)
	}

	var result []*dtos.RolePermissionDetailResponse
	for _, rp := range rolePermissions {
		response := &dtos.RolePermissionDetailResponse{
			ID:             rp.ID.String(),
			RoleID:         rp.RoleID,
			RoleName:       rp.RoleName,
			PermissionID:   rp.PermissionID,
			PermissionName: rp.PermissionName,
			Resource:       rp.Resource,
			Action:         rp.Action,
			Status:         model.UserStatusActive,
			CreatedAt:      utils.FormatTime(rp.CreatedAt.Time),
			UpdatedAt:      utils.FormatTime(rp.UpdatedAt.Time),
		}

		if rp.ScopeID.Valid {
			response.ScopeID = rp.ScopeID.String
		}
		if rp.ScopeName.Valid {
			response.ScopeName = rp.ScopeName.String
		}
		if rp.DeletedAt.Valid {
			response.DeletedAt = utils.FormatTime(rp.DeletedAt.Time)
		}

		result = append(result, response)
	}

	log.Info().
		Int("count", len(result)).
		Msg("Successfully retrieved role permissions")

	return result, nil
}

func (s *rolePermissionService) GetRolePermissionByID(ctx context.Context, id string) (*dtos.RolePermissionDetailResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RolePermissionService").
		Str("method", "GetRolePermissionByID").
		Str("role_permission_id", id).
		Str("user_id", userID).
		Msg("Getting role permission by ID")

	var uuid pgtype.UUID
	err = uuid.Scan(id)
	if err != nil {
		log.Error().Err(err).Str("role_permission_id", id).Msg("Invalid UUID format")
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	rolePermission, err := s.repo.GetRolePermissionByID(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Str("role_permission_id", id).Msg("Failed to get role permission from repository")
		return nil, fmt.Errorf("failed to get role permission from repository: %w", err)
	}

	response := &dtos.RolePermissionDetailResponse{
		ID:             rolePermission.ID.String(),
		RoleID:         rolePermission.RoleID,
		RoleName:       rolePermission.RoleName,
		PermissionID:   rolePermission.PermissionID,
		PermissionName: rolePermission.PermissionName,
		Resource:       rolePermission.Resource,
		Action:         rolePermission.Action,
		Status:         model.UserStatusActive,
		CreatedAt:      utils.FormatTime(rolePermission.CreatedAt.Time),
		UpdatedAt:      utils.FormatTime(rolePermission.UpdatedAt.Time),
	}

	if rolePermission.ScopeID.Valid {
		response.ScopeID = rolePermission.ScopeID.String
	}
	if rolePermission.ScopeName.Valid {
		response.ScopeName = rolePermission.ScopeName.String
	}
	if rolePermission.DeletedAt.Valid {
		response.DeletedAt = utils.FormatTime(rolePermission.DeletedAt.Time)
	}

	log.Info().
		Str("role_permission_id", id).
		Str("role_name", response.RoleName).
		Str("permission_name", response.PermissionName).
		Msg("Successfully retrieved role permission")

	return response, nil
}

func (s *rolePermissionService) CreateRolePermission(ctx context.Context, req *dtos.CreateRolePermissionRequest) (*dtos.RolePermissionResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RolePermissionService").
		Str("method", "CreateRolePermission").
		Str("role_id", req.RoleID).
		Str("permission_id", req.PermissionID).
		Str("user_id", userID).
		Msg("Creating role permission")

	params := repository.CreateRolePermissionParams{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
		Status:       repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	if req.ScopeID != "" {
		params.ScopeID = pgtype.Text{String: req.ScopeID, Valid: true}
	}

	rolePermission, err := s.repo.CreateRolePermission(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to create role permission in repository")
		return nil, fmt.Errorf("failed to create role permission in repository: %w", err)
	}

	response := &dtos.RolePermissionResponse{
		ID:           rolePermission.ID.String(),
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
		Status:       model.UserStatusActive,
		CreatedAt:    utils.FormatTime(rolePermission.CreatedAt.Time),
		UpdatedAt:    utils.FormatTime(rolePermission.UpdatedAt.Time),
	}

	if rolePermission.ScopeID.Valid {
		response.ScopeID = rolePermission.ScopeID.String
	}
	if rolePermission.DeletedAt.Valid {
		response.DeletedAt = utils.FormatTime(rolePermission.DeletedAt.Time)
	}

	log.Info().
		Str("role_permission_id", response.ID).
		Str("role_id", response.RoleID).
		Str("permission_id", response.PermissionID).
		Msg("Successfully created role permission")

	return response, nil
}
