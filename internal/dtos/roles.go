package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type Role struct {
	model.BaseModel
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsSystemRole bool   `json:"is_system_role"`
	DeletedAt    string `json:"deleted_at"`
}

type RoleResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsSystemRole bool   `json:"is_system_role"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type RolesListResponse struct {
	Roles []RoleResponse `json:"roles"`
	Meta  PaginationMeta `json:"meta"`
}

type CreateRoleRequest struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	IsSystemRole bool   `json:"is_system_role"`
	Status       string `json:"status"`
}

type UpdateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (r *Role) ToResponse() *RoleResponse {
	description := ""
	if r.Description != "" {
		description = r.Description
	}

	deletedAt := ""
	if r.DeletedAt != "" {
		deletedAt = r.DeletedAt
	}

	return &RoleResponse{
		ID:           r.ID,
		Name:         r.Name,
		Description:  description,
		IsSystemRole: r.IsSystemRole,
		Status:       r.Status.String,
		CreatedAt:    utils.FormatTime(r.CreatedAt.Time),
		UpdatedAt:    utils.FormatTime(r.UpdatedAt.Time),
		DeletedAt:    deletedAt,
	}
}

func (r *Role) FromRepositoryModel(repo repository.Role) *Role {
	description := ""
	if repo.Description.Valid {
		description = repo.Description.String
	}

	deletedAt := ""
	if repo.DeletedAt.Valid {
		deletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return &Role{
		BaseModel: model.BaseModel{
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		ID:           repo.ID,
		Name:         repo.Name,
		Description:  description,
		IsSystemRole: repo.IsSystemRole.Bool,
		DeletedAt:    deletedAt,
	}
}

func NewRolesListResponse(data []RoleResponse, page, pageSize int, total int64) *RolesListResponse {
	return &RolesListResponse{
		Roles: data,
		Meta:  CreatePaginationMeta(page, pageSize, total),
	}
}
