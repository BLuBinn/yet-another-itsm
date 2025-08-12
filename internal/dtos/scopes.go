package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type Scope struct {
	model.BaseModel
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DeletedAt   string `json:"deleted_at"`
}

type ScopeResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type ScopesListResponse struct {
	Scopes []ScopeResponse `json:"scopes"`
	Meta   PaginationMeta  `json:"meta"`
}

type CreateScopeRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateScopeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (s *Scope) ToResponse() *ScopeResponse {
	description := ""
	if s.Description != "" {
		description = s.Description
	}

	deletedAt := ""
	if s.DeletedAt != "" {
		deletedAt = s.DeletedAt
	}

	return &ScopeResponse{
		ID:          s.ID,
		Name:        s.Name,
		Description: description,
		Status:      s.Status.String,
		CreatedAt:   utils.FormatTime(s.CreatedAt.Time),
		UpdatedAt:   utils.FormatTime(s.UpdatedAt.Time),
		DeletedAt:   deletedAt,
	}
}

func (s *Scope) FromRepositoryModel(repo repository.Scope) *Scope {
	description := ""
	if repo.Description.Valid {
		description = repo.Description.String
	}

	deletedAt := ""
	if repo.DeletedAt.Valid {
		deletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return &Scope{
		BaseModel: model.BaseModel{
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		ID:          repo.ID,
		Name:        repo.Name,
		Description: description,
		DeletedAt:   deletedAt,
	}
}

func NewScopesListResponse(data []ScopeResponse, page, pageSize int, total int64) *ScopesListResponse {
	return &ScopesListResponse{
		Scopes: data,
		Meta:   CreatePaginationMeta(page, pageSize, total),
	}
}
