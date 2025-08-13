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

type FormCategoryService interface {
	GetFormCategories(ctx context.Context) ([]*dtos.FormCategory, error)
	GetFormCategoryByID(ctx context.Context, id string) (*dtos.FormCategory, error)
	CreateFormCategory(ctx context.Context, name string, description string) (*dtos.FormCategory, error)
	UpdateFormCategory(ctx context.Context, id string, name string, description string) (*dtos.FormCategory, error)
	DeleteFormCategory(ctx context.Context, id string) error
}

type formCategoryService struct {
	repo *repository.Queries
}

func NewFormCategoryService(repo *repository.Queries) FormCategoryService {
	return &formCategoryService{
		repo: repo,
	}
}

func (s *formCategoryService) GetFormCategories(ctx context.Context) ([]*dtos.FormCategory, error) {
	log.Info().
		Str("service", "FormCategoryService").
		Str("method", "GetFormCategories").
		Msg("Getting all form categories")

	categories, err := s.repo.GetFormCategories(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get form categories from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormCategories, err)
	}

	result := make([]*dtos.FormCategory, len(categories))
	for i, category := range categories {
		result[i] = &dtos.FormCategory{}
		result[i].FromRepositoryModel(category)
	}

	return result, nil
}

func (s *formCategoryService) GetFormCategoryByID(ctx context.Context, id string) (*dtos.FormCategory, error) {
	log.Info().
		Str("service", "FormCategoryService").
		Str("method", "GetFormCategoryByID").
		Str("id", id).
		Msg("Getting form category by ID")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid form category ID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	category, err := s.repo.GetFormCategoryByID(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to get form category from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormCategory, err)
	}

	result := &dtos.FormCategory{}
	result.FromRepositoryModel(category)

	return result, nil
}

func (s *formCategoryService) CreateFormCategory(ctx context.Context, name string, description string) (*dtos.FormCategory, error) {
	log.Info().Str("service", "FormCategoryService").Str("method", "CreateFormCategory").Msg("Creating form category")

	category, err := s.repo.CreateFormCategory(ctx, repository.CreateFormCategoryParams{
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToCreateFormCategory)
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToCreateFormCategory, err)
	}

	formCategory := &dtos.FormCategory{}
	result := formCategory.FromRepositoryModel(category)
	return &result, nil
}

func (s *formCategoryService) UpdateFormCategory(ctx context.Context, id string, name string, description string) (*dtos.FormCategory, error) {
	log.Info().Str("service", "FormCategoryService").Str("method", "UpdateFormCategory").Str("id", id).Msg("Updating form category")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrInvalidUUIDFormat)
		return nil, fmt.Errorf("%s: %w", constants.ErrInvalidUUIDFormat, err)
	}

	category, err := s.repo.UpdateFormCategory(ctx, repository.UpdateFormCategoryParams{
		ID:          pgtype.UUID{Bytes: uuid, Valid: true},
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToUpdateFormCategory)
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToUpdateFormCategory, err)
	}

	formCategory := &dtos.FormCategory{}
	result := formCategory.FromRepositoryModel(category)
	return &result, nil
}

func (s *formCategoryService) DeleteFormCategory(ctx context.Context, id string) error {
	log.Info().
		Str("service", "FormCategoryService").
		Str("method", "DeleteFormCategory").
		Str("id", id).
		Msg("Deleting form category")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid form category ID format")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	err = s.repo.DeleteFormCategory(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to delete form category in repository")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToDeleteFormCategory, err)
	}

	return nil
}
