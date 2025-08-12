package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type PermissionService interface {
	GetAllPermissions(ctx context.Context) ([]*dtos.Permission, error)
	GetPermissionByID(ctx context.Context, id string) (*dtos.Permission, error)
	GetPermissionsByResource(ctx context.Context, resource string) ([]*dtos.Permission, error)
	GetPermissionsByResourceAndAction(ctx context.Context, resource, action string) (*dtos.Permission, error)
	GetActivePermissions(ctx context.Context) ([]*dtos.Permission, error)
	CreatePermission(ctx context.Context, req *dtos.CreatePermissionRequest) (*dtos.Permission, error)
	UpdatePermission(ctx context.Context, id string, req *dtos.UpdatePermissionRequest) (*dtos.Permission, error)
	DeletePermission(ctx context.Context, id string) error
	ResourceExists(ctx context.Context, resource string) (bool, error)
	ActionExists(ctx context.Context, action string) (bool, error)
	ResourceAndActionExists(ctx context.Context, resource, action string) (bool, error)
}

type permissionService struct {
	repo *repository.Queries
}

func NewPermissionService(repo *repository.Queries) PermissionService {
	return &permissionService{
		repo: repo,
	}
}

func (s *permissionService) GetAllPermissions(ctx context.Context) ([]*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "GetAllPermissions").
		Str("user_id", userID).
		Msg("Getting all permissions")

	permissions, err := s.repo.GetAllPermissions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get permissions from repository")
		return nil, fmt.Errorf("failed to get permissions from repository: %w", err)
	}

	var result []*dtos.Permission
	for _, permission := range permissions {
		permissionDTO := &dtos.Permission{}
		result = append(result, permissionDTO.FromRepositoryModel(permission))
	}

	log.Info().
		Int("count", len(result)).
		Msg("Successfully retrieved permissions")

	return result, nil
}

func (s *permissionService) GetPermissionByID(ctx context.Context, id string) (*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "GetPermissionByID").
		Str("permission_id", id).
		Str("user_id", userID).
		Msg("Getting permission by ID")

	permission, err := s.repo.GetPermissionByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("permission_id", id).Msg("Failed to get permission from repository")
		return nil, fmt.Errorf("failed to get permission from repository: %w", err)
	}

	permissionDTO := &dtos.Permission{}
	result := permissionDTO.FromRepositoryModel(permission)

	log.Info().
		Str("permission_id", id).
		Str("permission_name", result.Name).
		Msg("Successfully retrieved permission")

	return result, nil
}

func (s *permissionService) GetPermissionsByResource(ctx context.Context, resource string) ([]*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "GetPermissionsByResource").
		Str("resource", resource).
		Str("user_id", userID).
		Msg("Getting permissions by resource")

	permissions, err := s.repo.GetPermissionsByResource(ctx, resource)
	if err != nil {
		log.Error().Err(err).Str("resource", resource).Msg("Failed to get permissions from repository")
		return nil, fmt.Errorf("failed to get permissions from repository: %w", err)
	}

	var result []*dtos.Permission
	for _, permission := range permissions {
		permissionDTO := &dtos.Permission{}
		result = append(result, permissionDTO.FromRepositoryModel(permission))
	}

	return result, nil
}

func (s *permissionService) GetPermissionsByResourceAndAction(ctx context.Context, resource, action string) (*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "GetPermissionsByResourceAndAction").
		Str("resource", resource).
		Str("action", action).
		Str("user_id", userID).
		Msg("Getting permission by resource and action")

	params := repository.GetPermissionsByResourceAndActionParams{
		Resource: resource,
		Action:   action,
	}

	permission, err := s.repo.GetPermissionsByResourceAndAction(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to get permission from repository")
		return nil, fmt.Errorf("failed to get permission from repository: %w", err)
	}

	permissionDTO := &dtos.Permission{}
	result := permissionDTO.FromRepositoryModel(permission)

	return result, nil
}

func (s *permissionService) GetActivePermissions(ctx context.Context) ([]*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "GetActivePermissions").
		Str("user_id", userID).
		Msg("Getting active permissions")

	permissions, err := s.repo.GetActivePermissions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get active permissions from repository")
		return nil, fmt.Errorf("failed to get active permissions from repository: %w", err)
	}

	var result []*dtos.Permission
	for _, permission := range permissions {
		permissionDTO := &dtos.Permission{}
		result = append(result, permissionDTO.FromRepositoryModel(permission))
	}

	return result, nil
}

func (s *permissionService) CreatePermission(ctx context.Context, req *dtos.CreatePermissionRequest) (*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "CreatePermission").
		Str("permission_id", req.ID).
		Str("user_id", userID).
		Msg("Creating permission")

	params := repository.CreatePermissionParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: true},
		Resource:    req.Resource,
		Action:      req.Action,
		Status:      repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	permission, err := s.repo.CreatePermission(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to create permission in repository")
		return nil, fmt.Errorf("failed to create permission in repository: %w", err)
	}

	permissionDTO := &dtos.Permission{}
	result := permissionDTO.FromRepositoryModel(permission)

	log.Info().
		Str("permission_id", result.ID).
		Str("permission_name", result.Name).
		Msg("Successfully created permission")

	return result, nil
}

func (s *permissionService) UpdatePermission(ctx context.Context, id string, req *dtos.UpdatePermissionRequest) (*dtos.Permission, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "UpdatePermission").
		Str("permission_id", id).
		Str("user_id", userID).
		Msg("Updating permission")

	params := repository.UpdatePermissionParams{
		ID:          id,
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: true},
		Resource:    req.Resource,
		Action:      req.Action,
		Status:      repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	permission, err := s.repo.UpdatePermission(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to update permission in repository")
		return nil, fmt.Errorf("failed to update permission in repository: %w", err)
	}

	permissionDTO := &dtos.Permission{}
	result := permissionDTO.FromRepositoryModel(permission)

	log.Info().
		Str("permission_id", result.ID).
		Str("permission_name", result.Name).
		Msg("Successfully updated permission")

	return result, nil
}

func (s *permissionService) DeletePermission(ctx context.Context, id string) error {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "PermissionService").
		Str("method", "DeletePermission").
		Str("permission_id", id).
		Str("user_id", userID).
		Msg("Deleting permission")

	err = s.repo.DeletePermission(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("permission_id", id).Msg("Failed to delete permission in repository")
		return fmt.Errorf("failed to delete permission in repository: %w", err)
	}

	log.Info().
		Str("permission_id", id).
		Msg("Successfully deleted permission")

	return nil
}

func (s *permissionService) ResourceExists(ctx context.Context, resource string) (bool, error) {
	permissions, err := s.repo.GetPermissionsByResource(ctx, resource)
	if err != nil {
		return false, err
	}
	return len(permissions) > 0, nil
}

func (s *permissionService) ActionExists(ctx context.Context, action string) (bool, error) {
	permissions, err := s.repo.GetAllPermissions(ctx)
	if err != nil {
		return false, err
	}
	for _, permission := range permissions {
		if permission.Action == action {
			return true, nil
		}
	}
	return false, nil
}

func (s *permissionService) ResourceAndActionExists(ctx context.Context, resource, action string) (bool, error) {
	params := repository.GetPermissionsByResourceAndActionParams{
		Resource: resource,
		Action:   action,
	}
	_, err := s.repo.GetPermissionsByResourceAndAction(ctx, params)
	if err != nil {
		return false, nil
	}
	return true, nil
}
