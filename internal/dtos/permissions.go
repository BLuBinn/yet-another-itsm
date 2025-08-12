package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type Permission struct {
	model.BaseModel
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	DeletedAt   string `json:"deleted_at"`
}

type PermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type PermissionsListResponse struct {
	Permissions []PermissionResponse `json:"permissions"`
	Meta        PaginationMeta       `json:"meta"`
}

type CreatePermissionRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Resource    string `json:"resource" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Status      string `json:"status"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Resource    string `json:"resource" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Status      string `json:"status"`
}

func (p *Permission) ToResponse() *PermissionResponse {
	description := ""
	if p.Description != "" {
		description = p.Description
	}

	deletedAt := ""
	if p.DeletedAt != "" {
		deletedAt = p.DeletedAt
	}

	return &PermissionResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: description,
		Resource:    p.Resource,
		Action:      p.Action,
		Status:      p.Status.String,
		CreatedAt:   utils.FormatTime(p.CreatedAt.Time),
		UpdatedAt:   utils.FormatTime(p.UpdatedAt.Time),
		DeletedAt:   deletedAt,
	}
}

func (p *Permission) FromRepositoryModel(repo repository.Permission) *Permission {
	description := ""
	if repo.Description.Valid {
		description = repo.Description.String
	}

	deletedAt := ""
	if repo.DeletedAt.Valid {
		deletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return &Permission{
		BaseModel: model.BaseModel{
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		ID:          repo.ID,
		Name:        repo.Name,
		Description: description,
		Resource:    repo.Resource,
		Action:      repo.Action,
		DeletedAt:   deletedAt,
	}
}

func NewPermissionsListResponse(data []PermissionResponse, page, pageSize int, total int64) *PermissionsListResponse {
	return &PermissionsListResponse{
		Permissions: data,
		Meta:        CreatePaginationMeta(page, pageSize, total),
	}
}
