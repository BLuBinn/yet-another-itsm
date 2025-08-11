package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type DepartmentService interface {
	GetDepartmentByID(ctx context.Context, id string) (*dtos.Department, error)
	GetDepartmentByName(ctx context.Context, name string, businessUnitID string) (*dtos.Department, error)
}

type departmentService struct {
	repo *repository.Queries
}

func NewDepartmentService(repo *repository.Queries) DepartmentService {
	return &departmentService{
		repo: repo,
	}
}

// GetDepartmentByID gets a department by ID.
func (s *departmentService) GetDepartmentByID(ctx context.Context, id string) (*dtos.Department, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	var uuid pgtype.UUID
	err = uuid.Scan(id)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	log.Info().
		Str("service", "DepartmentService").
		Str("endpoint", "GetDepartmentByID").
		Str("id", id).
		Str("user_id", userID).
		Msg("Getting department by ID")

	repoDepartment, err := s.repo.GetDepartmentByID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetDepartment, err)
	}

	dto := (&dtos.Department{}).FromRepositoryModel(repoDepartment)
	return dto, nil
}

// GetDepartmentByName gets a department by name.
func (s *departmentService) GetDepartmentByName(ctx context.Context, name string, businessUnitID string) (*dtos.Department, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	var uuid pgtype.UUID
	err = uuid.Scan(businessUnitID)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrInvalidBusinessUnitUUIDFormat, err)
	}

	log.Info().
		Str("service", "DepartmentService").
		Str("endpoint", "GetDepartmentByName").
		Str("name", name).
		Str("business_unit_id", businessUnitID).
		Str("user_id", userID).
		Msg("Getting department by name")

	params := repository.GetDepartmentByNameParams{
		Name:           name,
		BusinessUnitID: uuid,
	}

	repoDepartment, err := s.repo.GetDepartmentByName(ctx, params)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetDepartment, err)
	}

	dto := (&dtos.Department{}).FromRepositoryModel(repoDepartment)
	return dto, nil
}
