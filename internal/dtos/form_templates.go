package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type FormTemplate struct {
	model.BaseModel
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	FormCategoryID string `json:"form_category_id"`
	BusinessUnitID string `json:"business_unit_id"`
	Version        int32  `json:"version"`
	PublishedAt    string `json:"published_at"`
	CreatedBy      string `json:"created_by"`
	ApprovedBy     string `json:"approved_by"`
	ApprovedAt     string `json:"approved_at"`
	DeletedAt      string `json:"deleted_at"`
}

type FormTemplateResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	FormCategoryID string `json:"form_category_id"`
	BusinessUnitID string `json:"business_unit_id"`
	Version        int32  `json:"version"`
	PublishedAt    string `json:"published_at"`
	CreatedBy      string `json:"created_by"`
	ApprovedBy     string `json:"approved_by"`
	ApprovedAt     string `json:"approved_at"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type CreateFormTemplateRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	FormCategoryID string `json:"form_category_id" binding:"required"`
	BusinessUnitID string `json:"business_unit_id" binding:"required"`
	Version        int32  `json:"version"`
	CreatedBy      string `json:"created_by"`
}

type UpdateFormTemplateRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	FormCategoryID string `json:"form_category_id"`
	BusinessUnitID string `json:"business_unit_id"`
	Version        int32  `json:"version"`
}

type PublishFormTemplateRequest struct {
	ApprovedBy string `json:"approved_by" binding:"required"`
}

type FormTemplatesListResponse struct {
	Items      []FormTemplateResponse `json:"items"`
	Page       int                    `json:"page"`
	Size       int                    `json:"size"`
	TotalItems int64                  `json:"total_items"`
}

func NewFormTemplatesListResponse(items []FormTemplateResponse, page, size int, totalItems int64) *FormTemplatesListResponse {
	return &FormTemplatesListResponse{
		Items:      items,
		Page:       page,
		Size:       size,
		TotalItems: totalItems,
	}
}

func (ft *FormTemplate) ToResponse() *FormTemplateResponse {
	return &FormTemplateResponse{
		ID:             ft.ID,
		Name:           ft.Name,
		Description:    ft.Description,
		FormCategoryID: ft.FormCategoryID,
		BusinessUnitID: ft.BusinessUnitID,
		Version:        ft.Version,
		PublishedAt:    ft.PublishedAt,
		CreatedBy:      ft.CreatedBy,
		ApprovedBy:     ft.ApprovedBy,
		ApprovedAt:     ft.ApprovedAt,
		Status:         ft.Status.String,
		CreatedAt:      utils.FormatTime(ft.CreatedAt.Time),
		UpdatedAt:      utils.FormatTime(ft.UpdatedAt.Time),
		DeletedAt:      ft.DeletedAt,
	}
}

func (ft *FormTemplate) FromRepositoryModel(repo repository.FormTemplate) FormTemplate {
	template := FormTemplate{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		Name:           repo.Name,
		Description:    repo.Description.String,
		FormCategoryID: repo.FormCategoryID.String(),
		BusinessUnitID: repo.BusinessUnitID.String(),
		Version:        repo.Version.Int32,
		CreatedBy:      repo.CreatedBy.String(),
	}

	if repo.PublishedAt.Valid {
		template.PublishedAt = utils.FormatTime(repo.PublishedAt.Time)
	}
	if repo.DeletedAt.Valid {
		template.DeletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return template
}
