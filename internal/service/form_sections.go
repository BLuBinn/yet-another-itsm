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

type FormSectionService interface {
	GetFormSections(ctx context.Context, formTemplateID string) ([]*dtos.FormSection, error)
	GetFormSectionByID(ctx context.Context, id string) (*dtos.FormSection, error)
	CreateFormSection(ctx context.Context, req *dtos.CreateFormSectionRequest) (*dtos.FormSection, error)
	UpdateFormSection(ctx context.Context, id string, req *dtos.UpdateFormSectionRequest) (*dtos.FormSection, error)
	DeleteFormSection(ctx context.Context, id string) error
}

type formSectionService struct {
	repo *repository.Queries
}

func NewFormSectionService(repo *repository.Queries) FormSectionService {
	return &formSectionService{
		repo: repo,
	}
}

func (s *formSectionService) GetFormSections(ctx context.Context, formTemplateID string) ([]*dtos.FormSection, error) {
	log.Info().
		Str("service", "FormSectionService").
		Str("method", "GetFormSections").
		Msg("Getting all form sections")

	uuid, err := utils.ParseUUID(formTemplateID)
	if err != nil {
		log.Error().Err(err).Str("formTemplateID", formTemplateID).Msg("Invalid form template UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	sections, err := s.repo.GetFormSections(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get form sections from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormSections, err)
	}

	result := make([]*dtos.FormSection, len(sections))
	for i, section := range sections {
		result[i] = &dtos.FormSection{}
		result[i].FromRepositoryModel(section)
	}

	return result, nil
}

func (s *formSectionService) GetFormSectionByID(ctx context.Context, id string) (*dtos.FormSection, error) {
	log.Info().
		Str("service", "FormSectionService").
		Str("method", "GetFormSectionByID").
		Str("id", id).
		Msg("Getting form section by ID")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	section, err := s.repo.GetFormSectionByID(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to get form section from repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetFormSection, err)
	}

	result := &dtos.FormSection{}
	result.FromRepositoryModel(section)
	return result, nil
}

func (s *formSectionService) CreateFormSection(ctx context.Context, req *dtos.CreateFormSectionRequest) (*dtos.FormSection, error) {
	log.Info().
		Str("service", "FormSectionService").
		Str("method", "CreateFormSection").
		Str("name", req.SectionName).
		Str("templateID", req.FormTemplateID).
		Msg("Creating form section")

	templateUUID, err := utils.ParseUUID(req.FormTemplateID)
	if err != nil {
		log.Error().Err(err).Str("templateID", req.FormTemplateID).Msg("Invalid template UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	params := repository.CreateFormSectionParams{
		SectionName:    req.SectionName,
		Description:    pgtype.Text{String: req.Description, Valid: req.Description != ""},
		FormTemplateID: pgtype.UUID{Bytes: templateUUID, Valid: true},
		SectionOrder:   req.SectionOrder,
	}

	section, err := s.repo.CreateFormSection(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create form section in repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToCreateFormSection, err)
	}

	result := &dtos.FormSection{}
	result.FromRepositoryModel(section)
	return result, nil
}

// self
func (s *formSectionService) UpdateFormSection(ctx context.Context, id string, req *dtos.UpdateFormSectionRequest) (*dtos.FormSection, error) {
	log.Info().
		Str("service", "FormSectionService").
		Str("method", "UpdateFormSection").
		Str("id", id).
		Msg("Updating form section")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	params := repository.UpdateFormSectionParams{
		ID:           pgtype.UUID{Bytes: uuid, Valid: true},
		SectionName:  req.SectionName,
		Description:  pgtype.Text{String: req.Description, Valid: req.Description != ""},
		SectionOrder: req.SectionOrder,
	}

	section, err := s.repo.UpdateFormSection(ctx, params)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to update form section in repository")
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToUpdateFormSection, err)
	}

	result := &dtos.FormSection{}
	result.FromRepositoryModel(section)
	return result, nil
}

func (s *formSectionService) DeleteFormSection(ctx context.Context, id string) error {
	log.Info().
		Str("service", "FormSectionService").
		Str("method", "DeleteFormSection").
		Str("id", id).
		Msg("Deleting form section")

	uuid, err := utils.ParseUUID(id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Invalid UUID format")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	err = s.repo.DeleteFormSection(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to delete form section from repository")
		return fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToDeleteFormSection, err)
	}

	return nil
}
