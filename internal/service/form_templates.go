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

type FormTemplateService interface {
	GetFormTemplates(ctx context.Context) ([]*dtos.FormTemplate, error)
	GetFormTemplateByID(ctx context.Context, id string) (*dtos.FormTemplate, error)
	GetFormTemplatesByCategory(ctx context.Context, categoryID string) ([]*dtos.FormTemplate, error)
	CreateFormTemplate(ctx context.Context, req *dtos.CreateFormTemplateRequest) (*dtos.FormTemplate, error)
	UpdateFormTemplate(ctx context.Context, id string, req *dtos.UpdateFormTemplateRequest) (*dtos.FormTemplate, error)
	// PublishFormTemplate(ctx context.Context, id string, req *dtos.PublishFormTemplateRequest) (*dtos.FormTemplate, error)
	DeleteFormTemplate(ctx context.Context, id string) error
}

type formTemplateService struct {
	repo *repository.Queries
}

func NewFormTemplateService(repo *repository.Queries) FormTemplateService {
	return &formTemplateService{
		repo: repo,
	}
}

func (s *formTemplateService) GetFormTemplates(ctx context.Context) ([]*dtos.FormTemplate, error) {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "GetFormTemplates").
		Msg("Getting all form templates")

	templates, err := s.repo.GetFormTemplates(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get form templates from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormTemplates, err)
	}

	result := make([]*dtos.FormTemplate, len(templates))
	for i, template := range templates {
		result[i] = &dtos.FormTemplate{}
		result[i].FromRepositoryModel(template)
	}

	return result, nil
}

func (s *formTemplateService) GetFormTemplateByID(ctx context.Context, id string) (*dtos.FormTemplate, error) {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "GetFormTemplateByID").
		Str("id", id).
		Msg("Getting form template by ID")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	template, err := s.repo.GetFormTemplateByID(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to get form template from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormTemplate, err)
	}

	result := &dtos.FormTemplate{}
	result.FromRepositoryModel(template)
	return result, nil
}

func (s *formTemplateService) GetFormTemplatesByCategory(ctx context.Context, categoryID string) ([]*dtos.FormTemplate, error) {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "GetFormTemplatesByCategory").
		Str("categoryID", categoryID).
		Msg("Getting form templates by category")

	uuid, err := utils.ParseUUID(categoryID)
	if err != nil {
		log.Error().Err(err).Str("categoryID", categoryID).Msg("Invalid UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	templates, err := s.repo.GetFormTemplatesByCategory(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("categoryID", categoryID).Msg("Failed to get form templates by category from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormTemplates, err)
	}

	result := make([]*dtos.FormTemplate, len(templates))
	for i, template := range templates {
		result[i] = &dtos.FormTemplate{}
		result[i].FromRepositoryModel(template)
	}

	return result, nil
}

func (s *formTemplateService) CreateFormTemplate(ctx context.Context, req *dtos.CreateFormTemplateRequest) (*dtos.FormTemplate, error) {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "CreateFormTemplate").
		Str("name", req.Name).
		Msg("Creating form template")

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user ID from context")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	categoryUUID, err := utils.ParseUUID(req.FormCategoryID)
	if err != nil {
		log.Error().Err(err).Str("categoryID", req.FormCategoryID).Msg("Invalid category UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	businessUnitUUID, err := utils.ParseUUID(req.BusinessUnitID)
	if err != nil {
		log.Error().Err(err).Str("businessUnitID", req.BusinessUnitID).Msg("Invalid business unit UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	userUUID, err := utils.ParseUUID(userID)
	if err != nil {
		log.Error().Err(err).Str("userID", userID).Msg("Invalid user UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	params := repository.CreateFormTemplateParams{
		Name:           req.Name,
		Description:    pgtype.Text{String: req.Description, Valid: req.Description != ""},
		FormCategoryID: pgtype.UUID{Bytes: categoryUUID, Valid: true},
		BusinessUnitID: pgtype.UUID{Bytes: businessUnitUUID, Valid: true},
		CreatedBy:      pgtype.UUID{Bytes: userUUID, Valid: true},
	}

	template, err := s.repo.CreateFormTemplate(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create form template in repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToCreateFormTemplate, err)
	}

	result := &dtos.FormTemplate{}
	result.FromRepositoryModel(template)
	return result, nil
}

func (s *formTemplateService) UpdateFormTemplate(ctx context.Context, id string, req *dtos.UpdateFormTemplateRequest) (*dtos.FormTemplate, error) {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "UpdateFormTemplate").
		Str("id", id).
		Msg("Updating form template")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	categoryUUID, err := utils.ParseUUID(req.FormCategoryID)
	if req.FormCategoryID != "" {
		if err != nil {
			log.Error().Err(err).Str("categoryID", req.FormCategoryID).Msg("Invalid category UUID format")
			return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
		}
	}

	businessUnitUUID, err := utils.ParseUUID(req.BusinessUnitID)
	if req.BusinessUnitID != "" {
		if err != nil {
			log.Error().Err(err).Str("businessUnitID", req.BusinessUnitID).Msg("Invalid business unit UUID format")
			return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
		}
	}

	params := repository.UpdateFormTemplateParams{
		ID:             pgtype.UUID{Bytes: uuid, Valid: true},
		Name:           req.Name,
		Description:    pgtype.Text{String: req.Description, Valid: req.Description != ""},
		FormCategoryID: pgtype.UUID{Bytes: categoryUUID, Valid: true},
		BusinessUnitID: pgtype.UUID{Bytes: businessUnitUUID, Valid: true},
	}

	template, err := s.repo.UpdateFormTemplate(ctx, params)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to update form template in repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToUpdateFormTemplate, err)
	}

	result := &dtos.FormTemplate{}
	result.FromRepositoryModel(template)
	return result, nil
}

// func (s *formTemplateService) PublishFormTemplate(ctx context.Context, id string, req *dtos.PublishFormTemplateRequest) (*dtos.FormTemplate, error) {
// 	log.Info().
// 		Str("service", "FormTemplateService").
// 		Str("method", "PublishFormTemplate").
// 		Str("id", id).
// 		Msg("Publishing form template")

// 	uuid, err := utils.ParseUUID(id)
// 	if err != nil {
// 		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
// 		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
// 	}

// 	params := repository.PublishFormTemplateParams{
// 		ID:      pgtype.UUID{Bytes: uuid, Valid: true},
// 		Version: pgtype.Text{String: req.Version, Valid: req.Version != ""},
// 	}

// 	template, err := s.repo.PublishFormTemplate(ctx, params)
// 	if err != nil {
// 		log.Error().Err(err).Str("id", id).Msg("Failed to publish form template in repository")
// 		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToPublishFormTemplate, err)
// 	}

// 	result := &dtos.FormTemplate{}
// 	result.FromRepositoryModel(template)
// 	return result, nil
// }

func (s *formTemplateService) DeleteFormTemplate(ctx context.Context, id string) error {
	log.Info().
		Str("service", "FormTemplateService").
		Str("method", "DeleteFormTemplate").
		Str("id", id).
		Msg("Deleting form template")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	err = s.repo.DeleteFormTemplate(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to delete form template from repository")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToDeleteFormTemplate, err)
	}

	return nil
}
