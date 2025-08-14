package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type BusinessUnitService interface {
	GetAllBusinessUnitsInTenant(ctx context.Context) ([]*dtos.BusinessUnit, error)
	GetBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error)
	GetBusinessUnitByID(ctx context.Context, id string) (*dtos.BusinessUnit, error)
	CreateBusinessUnit(ctx context.Context, req *dtos.CreateBusinessUnitRequest) (*dtos.BusinessUnit, error)
	GetOrCreateBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error)
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
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetTenantID, err)
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetAllBusinessUnitsInTenant").
		Str("tenant_id", tenantID).
		Str("user_id", userID).
		Msg("Getting business units for tenant")

	repoBusinessUnits, err := s.repo.GetAllBusinessUnitsInTenant(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetBusinessUnits, err)
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
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetBusinessUnitByDomainName").
		Str("domain_name", domainName).
		Str("user_id", userID).
		Msg("Getting business unit by domain name")

	repoBusinessUnit, err := s.repo.GetBusinessUnitByDomainName(ctx, domainName)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetBusinessUnit, err)
	}

	dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoBusinessUnit)
	return dto, nil
}

// GetBusinessUnitByID gets a business unit by ID.
func (s *businessUnitService) GetBusinessUnitByID(ctx context.Context, id string) (*dtos.BusinessUnit, error) {
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
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetBusinessUnitByID").
		Str("id", id).
		Str("user_id", userID).
		Msg("Getting business unit by ID")

	repoBusinessUnit, err := s.repo.GetBusinessUnitByID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetBusinessUnit, err)
	}

	dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoBusinessUnit)
	return dto, nil
}

// CreateBusinessUnit creates a new business unit
func (s *businessUnitService) CreateBusinessUnit(ctx context.Context, req *dtos.CreateBusinessUnitRequest) (*dtos.BusinessUnit, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "CreateBusinessUnit").
		Str("domain_name", req.DomainName).
		Str("name", req.Name).
		Str("user_id", userID).
		Msg("Creating new business unit")

	params := repository.CreateBusinessUnitParams{
		DomainName: req.DomainName,
		TenantID:   req.TenantID,
		Name:       req.Name,
		Status:     repository.NullStatusEnum{StatusEnum: repository.StatusEnum(req.Status), Valid: true},
	}

	repoBusinessUnit, err := s.repo.CreateBusinessUnit(ctx, params)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToCreateBusinessUnit, err)
	}

	dto := (&dtos.BusinessUnit{}).FromRepositoryModel(repoBusinessUnit)
	return dto, nil
}

// GetOrCreateBusinessUnitByDomainName gets a business unit by domain name or creates it if it doesn't exist
func (s *businessUnitService) GetOrCreateBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error) {
	if domainName == "" {
		return nil, fmt.Errorf("domain name cannot be empty")
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	tenantID, err := utils.GetTenantID(ctx)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrFailedToGetTenantID, err)
	}

	log.Info().
		Str("service", "BusinessUnitService").
		Str("endpoint", "GetOrCreateBusinessUnitByDomainName").
		Str("domain_name", domainName).
		Str("user_id", userID).
		Msg("Getting or creating business unit by domain name")

	// Try to get existing business unit first
	businessUnit, err := s.GetBusinessUnitByDomainName(ctx, domainName)
	if err == nil {
		log.Info().
			Str("domain_name", domainName).
			Str("business_unit_id", businessUnit.ID).
			Str("business_unit_name", businessUnit.Name).
			Msg("Successfully found existing business unit")
		return businessUnit, nil
	}

	// If business unit doesn't exist, create it
	log.Info().
		Str("domain_name", domainName).
		Msg("Business unit not found, creating new business unit")

	// Generate a default name from domain
	defaultName := fmt.Sprintf("Business Unit - %s", domainName)

	createReq := &dtos.CreateBusinessUnitRequest{
		DomainName: domainName,
		TenantID:   tenantID,
		Name:       defaultName,
		Status:     model.UserStatusActive,
	}

	newBusinessUnit, createErr := s.CreateBusinessUnit(ctx, createReq)
	if createErr != nil {
		log.Error().Err(createErr).
			Str("domain_name", domainName).
			Msg("Failed to create business unit")
		return nil, fmt.Errorf(utils.ErrorFormat, "failed to create business unit", createErr)
	}

	log.Info().
		Str("domain_name", domainName).
		Str("business_unit_id", newBusinessUnit.ID).
		Str("business_unit_name", newBusinessUnit.Name).
		Msg("Successfully created new business unit")

	return newBusinessUnit, nil
}
