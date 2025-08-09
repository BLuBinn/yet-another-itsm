package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	model.BaseModel
	AzureAdObjectID string `json:"azure_ad_object_id"`
	HomeTenantID    string `json:"home_tenant_id"`
	DepartmentID    string `json:"department_id"`
	ManagerID       string `json:"manager_id"`
	Mail            string `json:"mail"`
	DisplayName     string `json:"display_name"`
	GivenName       string `json:"given_name"`
	SurName         string `json:"sur_name"`
	JobTitle        string `json:"job_title"`
	OfficeLocation  string `json:"office_location"`
	LastLogin       string `json:"last_login"`
	LockedUntil     string `json:"locked_until"`
	DeletedAt       string `json:"deleted_at"`
}

type UserResponse struct {
	ID              string `json:"id"`
	AzureAdObjectID string `json:"azure_ad_object_id"`
	HomeTenantID    string `json:"home_tenant_id"`
	DepartmentID    string `json:"department_id"`
	ManagerID       string `json:"manager_id"`
	Mail            string `json:"mail"`
	DisplayName     string `json:"display_name"`
	GivenName       string `json:"given_name"`
	SurName         string `json:"sur_name"`
	JobTitle        string `json:"job_title"`
	OfficeLocation  string `json:"office_location"`
	IsActive        bool   `json:"is_active"`
	LastLogin       string `json:"last_login"`
	LockedUntil     string `json:"locked_until"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

type UsersListResponse struct {
	Users []UserResponse `json:"users"`
	Meta  PaginationMeta `json:"meta"`
}

type CreateUserRequest struct {
	AzureAdObjectID string `json:"azure_ad_object_id" binding:"required"`
	HomeTenantID    string `json:"home_tenant_id" binding:"required"`
	DepartmentID    string `json:"department_id"`
	ManagerID       string `json:"manager_id"`
	Mail            string `json:"mail" binding:"required,email"`
	DisplayName     string `json:"display_name" binding:"required"`
	GivenName       string `json:"given_name"`
	SurName         string `json:"sur_name"`
	JobTitle        string `json:"job_title"`
	OfficeLocation  string `json:"office_location"`
	IsActive        bool   `json:"is_active"`
}

type UpdateUserRequest struct {
	DepartmentID   string `json:"department_id"`
	ManagerID      string `json:"manager_id"`
	Mail           string `json:"mail" binding:"required,email"`
	DisplayName    string `json:"display_name" binding:"required"`
	GivenName      string `json:"given_name"`
	SurName        string `json:"sur_name"`
	JobTitle       string `json:"job_title"`
	OfficeLocation string `json:"office_location"`
	IsActive       bool   `json:"is_active"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:              u.ID.String(),
		AzureAdObjectID: u.AzureAdObjectID,
		HomeTenantID:    u.HomeTenantID,
		DepartmentID:    u.DepartmentID,
		ManagerID:       u.ManagerID,
		Mail:            u.Mail,
		DisplayName:     u.DisplayName,
		GivenName:       u.GivenName,
		SurName:         u.SurName,
		JobTitle:        u.JobTitle,
		OfficeLocation:  u.OfficeLocation,
		IsActive:        u.IsActive.Bool,
		LastLogin:       u.LastLogin,
		LockedUntil:     u.LockedUntil,
		CreatedAt:       u.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       u.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		DeletedAt:       u.DeletedAt,
	}
}

func (u *User) FromRepositoryModel(repo repository.User) *User {
	user := &User{
		BaseModel: model.BaseModel{
			ID:        repo.ID,
			IsActive:  repo.IsActive,
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		AzureAdObjectID: repo.AzureAdObjectID,
		HomeTenantID:    repo.HomeTenantID.String(),
		Mail:            repo.Mail,
		DisplayName:     repo.DisplayName,
	}

	if repo.DepartmentID.Valid {
		user.DepartmentID = repo.DepartmentID.String()
	}
	if repo.ManagerID.Valid {
		user.ManagerID = repo.ManagerID.String()
	}
	if repo.GivenName.Valid {
		user.GivenName = repo.GivenName.String
	}
	if repo.SurName.Valid {
		user.SurName = repo.SurName.String
	}
	if repo.JobTitle.Valid {
		user.JobTitle = repo.JobTitle.String
	}
	if repo.OfficeLocation.Valid {
		user.OfficeLocation = repo.OfficeLocation.String
	}
	if repo.LastLogin.Valid {
		user.LastLogin = repo.LastLogin.Time.Format("2006-01-02T15:04:05Z07:00")
	}
	if repo.LockedUntil.Valid {
		user.LockedUntil = repo.LockedUntil.Time.Format("2006-01-02T15:04:05Z07:00")
	}
	if repo.DeletedAt.Valid {
		user.DeletedAt = repo.DeletedAt.Time.Format("2006-01-02T15:04:05Z07:00")
	}

	return user
}

func NewUsersListResponse(data []UserResponse, page, pageSize int, total int64) *UsersListResponse {
	return &UsersListResponse{
		Users: data,
		Meta:  CreatePaginationMeta(page, pageSize, total),
	}
}
