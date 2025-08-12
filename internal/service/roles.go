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

type RoleService interface {
	GetAllRoles(ctx context.Context) ([]*dtos.Role, error)
	GetRoleByID(ctx context.Context, id string) (*dtos.Role, error)
	GetSystemRoles(ctx context.Context) ([]*dtos.Role, error)
	CreateRole(ctx context.Context, req *dtos.CreateRoleRequest) (*dtos.Role, error)
}

type roleService struct {
	repo *repository.Queries
}

func NewRoleService(repo *repository.Queries) RoleService {
	return &roleService{
		repo: repo,
	}
}

func (s *roleService) GetAllRoles(ctx context.Context) ([]*dtos.Role, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RoleService").
		Str("method", "GetAllRoles").
		Str("user_id", userID).
		Msg("Getting all roles")

	roles, err := s.repo.GetAllRoles(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get roles from repository")
		return nil, fmt.Errorf("failed to get roles from repository: %w", err)
	}

	var result []*dtos.Role
	for _, role := range roles {
		roleDTO := &dtos.Role{}
		result = append(result, roleDTO.FromRepositoryModel(role))
	}

	log.Info().
		Int("count", len(result)).
		Msg("Successfully retrieved roles")

	return result, nil
}

func (s *roleService) GetRoleByID(ctx context.Context, id string) (*dtos.Role, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RoleService").
		Str("method", "GetRoleByID").
		Str("role_id", id).
		Str("user_id", userID).
		Msg("Getting role by ID")

	role, err := s.repo.GetRoleByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("role_id", id).Msg("Failed to get role from repository")
		return nil, fmt.Errorf("failed to get role from repository: %w", err)
	}

	roleDTO := &dtos.Role{}
	result := roleDTO.FromRepositoryModel(role)

	log.Info().
		Str("role_id", id).
		Str("role_name", result.Name).
		Msg("Successfully retrieved role")

	return result, nil
}

func (s *roleService) GetSystemRoles(ctx context.Context) ([]*dtos.Role, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RoleService").
		Str("method", "GetSystemRoles").
		Str("user_id", userID).
		Msg("Getting system roles")

	roles, err := s.repo.GetSystemRoles(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get system roles from repository")
		return nil, fmt.Errorf("failed to get system roles from repository: %w", err)
	}

	var result []*dtos.Role
	for _, role := range roles {
		roleDTO := &dtos.Role{}
		result = append(result, roleDTO.FromRepositoryModel(role))
	}

	return result, nil
}

func (s *roleService) CreateRole(ctx context.Context, req *dtos.CreateRoleRequest) (*dtos.Role, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "RoleService").
		Str("method", "CreateRole").
		Str("role_id", req.ID).
		Str("user_id", userID).
		Msg("Creating role")

	params := repository.CreateRoleParams{
		ID:           req.ID,
		Name:         req.Name,
		Description:  pgtype.Text{String: req.Description, Valid: true},
		IsSystemRole: pgtype.Bool{Bool: req.IsSystemRole, Valid: true},
		Status:       repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	role, err := s.repo.CreateRole(ctx, params)
	if err != nil {
		log.Error().Err(err).Interface("params", params).Msg("Failed to create role in repository")
		return nil, fmt.Errorf("failed to create role in repository: %w", err)
	}

	roleDTO := &dtos.Role{}
	result := roleDTO.FromRepositoryModel(role)

	log.Info().
		Str("role_id", result.ID).
		Str("role_name", result.Name).
		Msg("Successfully created role")

	return result, nil
}
