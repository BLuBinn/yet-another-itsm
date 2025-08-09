package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type Department struct {
	model.BaseModel
	BusinessUnitID string `json:"business_unit_id"`
	Name           string `json:"name"`
	DeletedAt      string `json:"deleted_at"`
}

type DepartmentResponse struct {
	ID             string `json:"id"`
	BusinessUnitID string `json:"business_unit_id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type DepartmentsListResponse struct {
	Departments []DepartmentResponse `json:"departments"`
	Meta        PaginationMeta       `json:"meta"`
}

func (d *Department) ToResponse() *DepartmentResponse {
	return &DepartmentResponse{
		ID:             d.ID.String(),
		BusinessUnitID: d.BusinessUnitID,
		Name:           d.Name,
		IsActive:       d.IsActive.Bool,
		CreatedAt:      d.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      d.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		DeletedAt:      d.DeletedAt,
	}
}

func (d *Department) FromRepositoryModel(repo repository.Department) *Department {
	return &Department{
		BaseModel: model.BaseModel{
			ID:        repo.ID,
			IsActive:  repo.IsActive,
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		BusinessUnitID: repo.BusinessUnitID.String(),
		Name:           repo.Name,
	}
}

func NewDepartmentsListResponse(data []DepartmentResponse, page, pageSize int, total int64) *DepartmentsListResponse {
	return &DepartmentsListResponse{
		Departments: data,
		Meta:        CreatePaginationMeta(page, pageSize, total),
	}
}
