package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type FormCategory struct {
	model.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	DeletedAt   string `json:"deleted_at"`
}

type FormCategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type CreateFormCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateFormCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type FormCategoriesListResponse struct {
	Items      []FormCategoryResponse `json:"items"`
	Page       int                    `json:"page"`
	Size       int                    `json:"size"`
	TotalItems int64                  `json:"total_items"`
}

func NewFormCategoriesListResponse(items []FormCategoryResponse, page, size int, totalItems int64) *FormCategoriesListResponse {
	return &FormCategoriesListResponse{
		Items:      items,
		Page:       page,
		Size:       size,
		TotalItems: totalItems,
	}
}

func (fc *FormCategory) ToResponse() *FormCategoryResponse {
	return &FormCategoryResponse{
		ID:          fc.ID,
		Name:        fc.Name,
		Description: fc.Description,
		Status:      fc.Status.String,
		CreatedAt:   utils.FormatTime(fc.CreatedAt.Time),
		UpdatedAt:   utils.FormatTime(fc.UpdatedAt.Time),
		DeletedAt:   fc.DeletedAt,
	}
}

func (fc *FormCategory) FromRepositoryModel(repo repository.FormCategory) FormCategory {
	category := FormCategory{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		Name:        repo.Name,
		Description: repo.Description.String,
	}
	if repo.DeletedAt.Valid {
		category.DeletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}
	return category
}
