package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type RoleAssignment struct {
	model.BaseModel
	ID                pgtype.UUID `json:"id"`
	RolePermissionsID string      `json:"role_permissions_id"`
	AssigneeID        string      `json:"assignee_id"`
	BusinessUnitID    string      `json:"business_unit_id"`
	DepartmentID      string      `json:"department_id"`
	AssignedBy        string      `json:"assigned_by"`
	AssignedAt        string      `json:"assigned_at"`
	ExpiresAt         string      `json:"expires_at"`
	DeletedAt         string      `json:"deleted_at"`
	Status            string      `json:"status"`
	UpdatedAt         string      `json:"updated_at"`
	CreatedAt         string      `json:"created_at"`
	RoleName          string      `json:"role_name"`
	PermissionName    string      `json:"permission_name"`
	Resource          string      `json:"resource"`
	Action            string      `json:"action"`
	ScopeName         string      `json:"scope_name"`
	BusinessUnitName  string      `json:"business_unit_name"`
	DepartmentName    string      `json:"department_name"`
	AssigneeName      string      `json:"assignee_name"`
	AssigneeEmail     string      `json:"assignee_email"`
}

type RoleAssignmentResponse struct {
	ID                string `json:"id"`
	RolePermissionsID string `json:"role_permissions_id"`
	AssigneeID        string `json:"assignee_id"`
	BusinessUnitID    string `json:"business_unit_id"`
	DepartmentID      string `json:"department_id"`
	AssignedBy        string `json:"assigned_by"`
	AssignedAt        string `json:"assigned_at"`
	ExpiresAt         string `json:"expires_at"`
	Status            string `json:"status"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
}

type RoleAssignmentDetailResponse struct {
	ID                string `json:"id"`
	RolePermissionsID string `json:"role_permissions_id"`
	AssigneeID        string `json:"assignee_id"`
	AssigneeName      string `json:"assignee_name"`
	AssigneeEmail     string `json:"assignee_email"`
	BusinessUnitID    string `json:"business_unit_id"`
	BusinessUnitName  string `json:"business_unit_name"`
	DepartmentID      string `json:"department_id"`
	DepartmentName    string `json:"department_name"`
	AssignedBy        string `json:"assigned_by"`
	AssignedAt        string `json:"assigned_at"`
	ExpiresAt         string `json:"expires_at"`
	RoleName          string `json:"role_name"`
	PermissionName    string `json:"permission_name"`
	Resource          string `json:"resource"`
	Action            string `json:"action"`
	ScopeName         string `json:"scope_name"`
	Status            string `json:"status"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
}

type RoleAssignmentsListResponse struct {
	RoleAssignments []RoleAssignmentDetailResponse `json:"role_assignments"`
	Meta            PaginationMeta                 `json:"meta"`
}

type CreateRoleAssignmentRequest struct {
	RolePermissionsID string `json:"role_permissions_id" binding:"required"`
	AssigneeID        string `json:"assignee_id" binding:"required"`
	BusinessUnitID    string `json:"business_unit_id"`
	DepartmentID      string `json:"department_id"`
	AssignedBy        string `json:"assigned_by"`
	ExpiresAt         string `json:"expires_at"`
	Status            string `json:"status"`
}

type UpdateRoleAssignmentRequest struct {
	ExpiresAt string `json:"expires_at"`
	Status    string `json:"status"`
}

func (ra *RoleAssignment) ToResponse() *RoleAssignmentResponse {
	businessUnitID := ""
	if ra.BusinessUnitID != "" {
		businessUnitID = ra.BusinessUnitID
	}

	departmentID := ""
	if ra.DepartmentID != "" {
		departmentID = ra.DepartmentID
	}

	assignedBy := ""
	if ra.AssignedBy != "" {
		assignedBy = ra.AssignedBy
	}

	expiresAt := ""
	if ra.ExpiresAt != "" {
		expiresAt = ra.ExpiresAt
	}

	deletedAt := ""
	if ra.DeletedAt != "" {
		deletedAt = ra.DeletedAt
	}

	return &RoleAssignmentResponse{
		ID:                ra.ID.String(),
		RolePermissionsID: ra.RolePermissionsID,
		AssigneeID:        ra.AssigneeID,
		BusinessUnitID:    businessUnitID,
		DepartmentID:      departmentID,
		AssignedBy:        assignedBy,
		AssignedAt:        ra.AssignedAt,
		ExpiresAt:         expiresAt,
		Status:            ra.Status,
		UpdatedAt:         ra.UpdatedAt,
		DeletedAt:         deletedAt,
	}
}

func (ra *RoleAssignment) FromRepositoryModel(repo repository.RoleAssignment) *RoleAssignment {
	businessUnitID := ""
	if repo.BusinessUnitID.Valid {
		businessUnitID = repo.BusinessUnitID.String()
	}

	departmentID := ""
	if repo.DepartmentID.Valid {
		departmentID = repo.DepartmentID.String()
	}

	assignedBy := ""
	if repo.AssignedBy.Valid {
		assignedBy = repo.AssignedBy.String()
	}

	assignedAt := ""
	if repo.AssignedAt.Valid {
		assignedAt = utils.FormatTime(repo.AssignedAt.Time)
	}

	expiresAt := ""
	if repo.ExpiresAt.Valid {
		expiresAt = utils.FormatTime(repo.ExpiresAt.Time)
	}

	deletedAt := ""
	if repo.DeletedAt.Valid {
		deletedAt = utils.FormatTime(repo.DeletedAt.Time)
	}

	return &RoleAssignment{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		RolePermissionsID: repo.RolePermissionsID.String(),
		AssigneeID:        repo.AssigneeID.String(),
		BusinessUnitID:    businessUnitID,
		DepartmentID:      departmentID,
		AssignedBy:        assignedBy,
		AssignedAt:        assignedAt,
		ExpiresAt:         expiresAt,
		DeletedAt:         deletedAt,
	}
}

func NewRoleAssignmentsListResponse(data []RoleAssignmentDetailResponse, page, pageSize int, total int64) *RoleAssignmentsListResponse {
	return &RoleAssignmentsListResponse{
		RoleAssignments: data,
		Meta:            CreatePaginationMeta(page, pageSize, total),
	}
}
