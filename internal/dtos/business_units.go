package dtos

import (
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

type BusinessUnit struct {
	model.BaseModel
	DomainName string `json:"domain_name"`
	TenantID   string `json:"tenant_id"`
	Name       string `json:"name"`
	DeletedAt  string `json:"deleted_at"`
}

type BusinessUnitResponse struct {
	ID         string `json:"id"`
	DomainName string `json:"domain_name"`
	TenantID   string `json:"tenant_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type BusinessUnitsListResponse struct {
	BusinessUnits []BusinessUnitResponse `json:"business_units"`
	Meta          PaginationMeta         `json:"meta"`
}

func (b *BusinessUnit) ToResponse() *BusinessUnitResponse {
	return &BusinessUnitResponse{
		ID:         b.ID,
		DomainName: b.DomainName,
		TenantID:   b.TenantID,
		Name:       b.Name,
		Status:     b.Status.String,
		CreatedAt:  utils.FormatTime(b.CreatedAt.Time),
		UpdatedAt:  utils.FormatTime(b.UpdatedAt.Time),
		DeletedAt:  b.DeletedAt,
	}
}

func (b *BusinessUnit) FromRepositoryModel(repo repository.BusinessUnit) *BusinessUnit {
	return &BusinessUnit{
		BaseModel: model.BaseModel{
			ID:        repo.ID.String(),
			Status:    pgtype.Text{String: string(repo.Status.StatusEnum), Valid: repo.Status.Valid},
			CreatedAt: pgtype.Timestamptz{Time: repo.CreatedAt.Time, Valid: repo.CreatedAt.Valid},
			UpdatedAt: pgtype.Timestamptz{Time: repo.UpdatedAt.Time, Valid: repo.UpdatedAt.Valid},
		},
		DomainName: repo.DomainName,
		TenantID:   repo.TenantID,
		Name:       repo.Name,
	}
}

func NewBusinessUnitsListResponse(data []BusinessUnitResponse, page, pageSize int, total int64) *BusinessUnitsListResponse {
	return &BusinessUnitsListResponse{
		BusinessUnits: data,
		Meta:          CreatePaginationMeta(page, pageSize, total),
	}
}
