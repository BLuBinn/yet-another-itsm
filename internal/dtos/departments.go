package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type Department struct {
	model.BaseModel
	Name           string `json:"name"`
	DeletedAt      string `json:"deleted_at"`
}

type CreateDepartmentRequest struct {
	Name   string `json:"name" validate:"required,min=1,max=255"`
	Status string `json:"status,omitempty"`
}

type DepartmentResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
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
		ID:             d.ID,
		Name:           d.Name,
		Status:         d.Status.String,
		CreatedAt:      utils.FormatTime(d.CreatedAt.Time),
		UpdatedAt:      utils.FormatTime(d.UpdatedAt.Time),
		DeletedAt:      d.DeletedAt,
	}
}

func (d *Department) FromRepositoryModel(repo repository.Department) *Department {
	return &Department{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		Name:           repo.Name,
	}
}

func NewDepartmentsListResponse(data []DepartmentResponse, page, pageSize int, total int64) *DepartmentsListResponse {
	return &DepartmentsListResponse{
		Departments: data,
		Meta:        CreatePaginationMeta(page, pageSize, total),
	}
}
