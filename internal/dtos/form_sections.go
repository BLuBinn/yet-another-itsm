package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type FormSection struct {
	model.BaseModel
	ID               string `json:"id"`
	FormTemplateID   string `json:"form_template_id"`
	SectionName      string `json:"section_name"`
	SectionOrder     int32  `json:"section_order"`
	Description      string `json:"description"`
	ConditionalLogic string `json:"conditional_logic"`
	DeletedAt        string `json:"deleted_at"`
}

type FormSectionResponse struct {
	ID               string `json:"id"`
	FormTemplateID   string `json:"form_template_id"`
	SectionName      string `json:"section_name"`
	SectionOrder     int32  `json:"section_order"`
	Description      string `json:"description"`
	ConditionalLogic string `json:"conditional_logic"`
	Status           string `json:"status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
}

type CreateFormSectionRequest struct {
	FormTemplateID   string `json:"form_template_id" binding:"required"`
	SectionName      string `json:"section_name" binding:"required"`
	SectionOrder     int32  `json:"section_order" binding:"required"`
	Description      string `json:"description"`
	ConditionalLogic string `json:"conditional_logic"`
}

type UpdateFormSectionRequest struct {
	ID               string `json:"id" binding:"required"`
	SectionName      string `json:"section_name" binding:"required"`
	SectionOrder     int32  `json:"section_order"`
	Description      string `json:"description"`
	ConditionalLogic string `json:"conditional_logic"`
}

type FormSectionsListResponse struct {
	Items      []FormSectionResponse `json:"items"`
	Page       int                   `json:"page"`
	Size       int                   `json:"size"`
	TotalItems int64                 `json:"total_items"`
}

func NewFormSectionsListResponse(items []FormSectionResponse, page, size int, totalItems int64) *FormSectionsListResponse {
	return &FormSectionsListResponse{
		Items:      items,
		Page:       page,
		Size:       size,
		TotalItems: totalItems,
	}
}

func (fs *FormSection) ToResponse() *FormSectionResponse {
	return &FormSectionResponse{
		ID:               fs.ID,
		FormTemplateID:   fs.FormTemplateID,
		SectionName:      fs.SectionName,
		SectionOrder:     fs.SectionOrder,
		Description:      fs.Description,
		ConditionalLogic: fs.ConditionalLogic,
		Status:           fs.Status.String,
		CreatedAt:        utils.FormatTime(fs.CreatedAt.Time),
		UpdatedAt:        utils.FormatTime(fs.UpdatedAt.Time),
		DeletedAt:        fs.DeletedAt,
	}
}

func (fs *FormSection) FromRepositoryModel(repo repository.FormSection) FormSection {
	section := FormSection{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		ID:               repo.ID.String(),
		FormTemplateID:   repo.FormTemplateID.String(),
		SectionName:      repo.SectionName,
		SectionOrder:     repo.SectionOrder,
		Description:      repo.Description.String,
		ConditionalLogic: string(repo.ConditionalLogic),
	}

	if repo.DeletedAt.Valid {
		section.DeletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return section
}
