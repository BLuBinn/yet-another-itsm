package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type RolePermission struct {
	model.BaseModel
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
	ScopeID      string `json:"scope_id"`
	DeletedAt    string `json:"deleted_at"`
}

type RolePermissionResponse struct {
	ID           string `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
	ScopeID      string `json:"scope_id"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type RolePermissionDetailResponse struct {
	ID             string `json:"id"`
	RoleID         string `json:"role_id"`
	RoleName       string `json:"role_name"`
	PermissionID   string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
	Resource       string `json:"resource"`
	Action         string `json:"action"`
	ScopeID        string `json:"scope_id"`
	ScopeName      string `json:"scope_name"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type RolePermissionsListResponse struct {
	RolePermissions []RolePermissionDetailResponse `json:"role_permissions"`
	Meta            PaginationMeta                 `json:"meta"`
}

type CreateRolePermissionRequest struct {
	RoleID       string `json:"role_id" binding:"required"`
	PermissionID string `json:"permission_id" binding:"required"`
	ScopeID      string `json:"scope_id"`
	Status       string `json:"status"`
}

type UpdateRolePermissionRequest struct {
	ScopeID string `json:"scope_id"`
	Status  string `json:"status"`
}

func (rp *RolePermission) ToResponse() *RolePermissionResponse {
	scopeID := ""
	if rp.ScopeID != "" {
		scopeID = rp.ScopeID
	}

	deletedAt := ""
	if rp.DeletedAt != "" {
		deletedAt = rp.DeletedAt
	}

	return &RolePermissionResponse{
		ID:           rp.ID,
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
		ScopeID:      scopeID,
		Status:       rp.Status.String,
		CreatedAt:    utils.FormatTime(rp.CreatedAt.Time),
		UpdatedAt:    utils.FormatTime(rp.UpdatedAt.Time),
		DeletedAt:    deletedAt,
	}
}

func (rp *RolePermission) FromRepositoryModel(repo repository.RolePermission) *RolePermission {
	scopeID := ""
	if repo.ScopeID.Valid {
		scopeID = repo.ScopeID.String
	}

	deletedAt := ""
	if repo.DeletedAt.Valid {
		deletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return &RolePermission{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		RoleID:       repo.RoleID,
		PermissionID: repo.PermissionID,
		ScopeID:      scopeID,
		DeletedAt:    deletedAt,
	}
}

func NewRolePermissionsListResponse(data []RolePermissionDetailResponse, page, pageSize int, total int64) *RolePermissionsListResponse {
	return &RolePermissionsListResponse{
		RolePermissions: data,
		Meta:            CreatePaginationMeta(page, pageSize, total),
	}
}
