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

type BusinessUnitService interface {
	GetAllBusinessUnitsInTenant(ctx context.Context) ([]*dtos.BusinessUnit, error)
	GetBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error)
	GetBusinessUnitByID(ctx context.Context, id string) (*dtos.BusinessUnit, error)
}

type businessUnitService struct {
	repo *repository.Queries
}

func NewBusinessUnitService(repo *repository.Queries) BusinessUnitService {
	return &businessUnitService{
		repo: repo,
	}
}

// GetAllBusinessUnitsInTenant gets all business units in a tenant.
func (s *businessUnitService) GetAllBusinessUnitsInTenant(ctx context.Context) ([]*dtos.BusinessUnit, error) {
	tenantID, err := utils.GetTenantID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetTenantID, err)
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetAllBusinessUnitsInTenant").
		Str("tenant_id", tenantID).
		Str("user_id", userID).
		Msg("Getting business units for tenant")

	repoBusinessUnits, err := s.repo.GetAllBusinessUnitsInTenant(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetBusinessUnits, err)
	}

	var businessUnits []*dtos.BusinessUnit
	for _, repoUnit := range repoBusinessUnits {
		dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoUnit)
		businessUnits = append(businessUnits, dto)
	}

	return businessUnits, nil
}

// GetBusinessUnitByDomainName gets a business unit by domain name.
func (s *businessUnitService) GetBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetBusinessUnitByDomainName").
		Str("domain_name", domainName).
		Str("user_id", userID).
		Msg("Getting business unit by domain name")

	repoBusinessUnit, err := s.repo.GetBusinessUnitByDomainName(ctx, domainName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetBusinessUnit, err)
	}

	dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoBusinessUnit)
	return dto, nil
}

// GetBusinessUnitByID gets a business unit by ID.
func (s *businessUnitService) GetBusinessUnitByID(ctx context.Context, id string) (*dtos.BusinessUnit, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetUserID, err)
	}

	var uuid pgtype.UUID
	err = uuid.Scan(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrInvalidUUIDFormat, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetBusinessUnitByID").
		Str("id", id).
		Str("user_id", userID).
		Msg("Getting business unit by ID")

	repoBusinessUnit, err := s.repo.GetBusinessUnitByID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrFailedToGetBusinessUnit, err)
	}

	dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoBusinessUnit)
	return dto, nil
}
