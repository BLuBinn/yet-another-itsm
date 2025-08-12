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

type ScopeService interface {
	GetAllScopes(ctx context.Context) ([]*dtos.Scope, error)
	GetScopeByID(ctx context.Context, id string) (*dtos.Scope, error)
	CreateScope(ctx context.Context, req *dtos.CreateScopeRequest) (*dtos.Scope, error)
}

type scopeService struct {
	repo *repository.Queries
}

func NewScopeService(repo *repository.Queries) ScopeService {
	return &scopeService{
		repo: repo,
	}
}

func (s *scopeService) GetAllScopes(ctx context.Context) ([]*dtos.Scope, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "ScopeService").
		Str("method", "GetAllScopes").
		Str("user_id", userID).
		Msg("Getting all scopes")

	scopes, err := s.repo.GetAllScopes(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get scopes from repository")
		return nil, fmt.Errorf("failed to get scopes from repository: %w", err)
	}

	var result []*dtos.Scope
	for _, scope := range scopes {
		scopeDTO := &dtos.Scope{}
		result = append(result, scopeDTO.FromRepositoryModel(scope))
	}

	log.Info().
		Int("count", len(result)).
		Msg("Successfully retrieved scopes")

	return result, nil
}

func (s *scopeService) GetScopeByID(ctx context.Context, id string) (*dtos.Scope, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "ScopeService").
		Str("method", "GetScopeByID").
		Str("scope_id", id).
		Str("user_id", userID).
		Msg("Getting scope by ID")

	scope, err := s.repo.GetScopeByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("scope_id", id).Msg("Failed to get scope from repository")
		return nil, fmt.Errorf("failed to get scope from repository: %w", err)
	}

	scopeDTO := &dtos.Scope{}
	result := scopeDTO.FromRepositoryModel(scope)

	log.Info().
		Str("scope_id", id).
		Str("scope_name", result.Name).
		Msg("Successfully retrieved scope")

	return result, nil
}

func (s *scopeService) CreateScope(ctx context.Context, req *dtos.CreateScopeRequest) (*dtos.Scope, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "ScopeService").
		Str("method", "CreateScope").
		Str("user_id", userID).
		Str("scope_name", req.Name).
		Msg("Creating scope")

	params := repository.CreateScopeParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: true},
		Status:      repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	scope, err := s.repo.CreateScope(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create scope in repository")
		return nil, fmt.Errorf("failed to create scope in repository: %w", err)
	}

	scopeDTO := &dtos.Scope{}
	result := scopeDTO.FromRepositoryModel(scope)

	log.Info().
		Str("scope_id", result.ID).
		Str("scope_name", result.Name).
		Msg("Successfully created scope")

	return result, nil
}
