package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueries implements all required methods from repository.Querier
type MockQueries struct {
	mock.Mock
}

// Business Unit methods
func (m *MockQueries) GetAllBusinessUnitsInTenant(ctx context.Context, tenantID string) ([]repository.BusinessUnit, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0).([]repository.BusinessUnit), args.Error(1)
}

func (m *MockQueries) GetBusinessUnitByDomainName(ctx context.Context, domainName string) (repository.BusinessUnit, error) {
	args := m.Called(ctx, domainName)
	return args.Get(0).(repository.BusinessUnit), args.Error(1)
}

func (m *MockQueries) GetBusinessUnitByID(ctx context.Context, id pgtype.UUID) (repository.BusinessUnit, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.BusinessUnit), args.Error(1)
}

// Department methods
func (m *MockQueries) GetAllDepartmentsInBusinessUnit(ctx context.Context, businessUnitID pgtype.UUID) ([]repository.Department, error) {
	args := m.Called(ctx, businessUnitID)
	return args.Get(0).([]repository.Department), args.Error(1)
}

func (m *MockQueries) GetDepartmentByID(ctx context.Context, id pgtype.UUID) (repository.Department, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.Department), args.Error(1)
}

func (m *MockQueries) GetDepartmentByName(ctx context.Context, arg repository.GetDepartmentByNameParams) (repository.Department, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.Department), args.Error(1)
}

// User methods
func (m *MockQueries) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.User), args.Error(1)
}

func (m *MockQueries) GetAllUsersInDepartment(ctx context.Context, departmentID pgtype.UUID) ([]repository.User, error) {
	args := m.Called(ctx, departmentID)
	return args.Get(0).([]repository.User), args.Error(1)
}

func (m *MockQueries) GetUserByEmail(ctx context.Context, mail string) (repository.User, error) {
	args := m.Called(ctx, mail)
	return args.Get(0).(repository.User), args.Error(1)
}

func (m *MockQueries) GetUserByID(ctx context.Context, id pgtype.UUID) (repository.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.User), args.Error(1)
}

func TestNewBusinessUnitService(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	assert.NotNil(t, service)
	assert.IsType(t, &businessUnitService{}, service)
}

func newBusinessUnitServiceWithMock(querier repository.Querier) BusinessUnitService {
	return &businessUnitService{
		repo: &repository.Queries{}, // This needs to be adapted based on actual structure
	}
}

func TestBusinessUnitService_GetAllBusinessUnitsInTenant_Success(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	// Create context with correct keys
	ctx := context.WithValue(context.Background(), utils.TenantIDKey, "test-tenant")
	ctx = context.WithValue(ctx, utils.UserIDKey, "test-user")

	validTime := time.Now()
	validUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}
	validTimePg := pgtype.Timestamptz{Time: validTime, Valid: true}

	mockBusinessUnits := []repository.BusinessUnit{
		{
			ID:         validUUID,
			DomainName: "test.com",
			TenantID:   "test-tenant",
			Name:       "Test Business Unit",
			IsActive:   pgtype.Bool{Bool: true, Valid: true},
			CreatedAt:  validTimePg,
			UpdatedAt:  validTimePg,
		},
	}

	mockRepo.On("GetAllBusinessUnitsInTenant", ctx, "test-tenant").Return(mockBusinessUnits, nil)

	result, err := service.GetAllBusinessUnitsInTenant(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test.com", result[0].DomainName)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetAllBusinessUnitsInTenant_NoTenantID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.Background()

	_, err := service.GetAllBusinessUnitsInTenant(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetTenantID)
}

func TestBusinessUnitService_GetAllBusinessUnitsInTenant_NoUserID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.TenantIDKey, "test-tenant")

	_, err := service.GetAllBusinessUnitsInTenant(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetUserID)
}

func TestBusinessUnitService_GetAllBusinessUnitsInTenant_RepositoryError(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.TenantIDKey, "test-tenant")
	ctx = context.WithValue(ctx, utils.UserIDKey, "test-user")

	mockRepo.On("GetAllBusinessUnitsInTenant", ctx, "test-tenant").Return([]repository.BusinessUnit{}, errors.New("database error"))

	_, err := service.GetAllBusinessUnitsInTenant(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetBusinessUnits)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetBusinessUnitByDomainName_Success(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")

	validTime := time.Now()
	validUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}
	validTimePg := pgtype.Timestamptz{Time: validTime, Valid: true}

	mockBusinessUnit := repository.BusinessUnit{
		ID:         validUUID,
		DomainName: "test.com",
		TenantID:   "test-tenant",
		Name:       "Test Business Unit",
		IsActive:   pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:  validTimePg,
		UpdatedAt:  validTimePg,
	}

	mockRepo.On("GetBusinessUnitByDomainName", ctx, "test.com").Return(mockBusinessUnit, nil)

	result, err := service.GetBusinessUnitByDomainName(ctx, "test.com")

	assert.NoError(t, err)
	assert.Equal(t, "test.com", result.DomainName)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetBusinessUnitByDomainName_NoUserID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.Background()

	_, err := service.GetBusinessUnitByDomainName(ctx, "test.com")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetUserID)
}

func TestBusinessUnitService_GetBusinessUnitByDomainName_RepositoryError(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")

	mockRepo.On("GetBusinessUnitByDomainName", ctx, "test.com").Return(repository.BusinessUnit{}, errors.New("not found"))

	_, err := service.GetBusinessUnitByDomainName(ctx, "test.com")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetBusinessUnit)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetBusinessUnitByID_Success(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	validTime := time.Now()
	validUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}
	validTimePg := pgtype.Timestamptz{Time: validTime, Valid: true}

	mockBusinessUnit := repository.BusinessUnit{
		ID:         validUUID,
		DomainName: "test.com",
		TenantID:   "test-tenant",
		Name:       "Test Business Unit",
		IsActive:   pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:  validTimePg,
		UpdatedAt:  validTimePg,
	}

	mockRepo.On("GetBusinessUnitByID", ctx, mock.AnythingOfType("pgtype.UUID")).Return(mockBusinessUnit, nil)

	result, err := service.GetBusinessUnitByID(ctx, uuidStr)

	assert.NoError(t, err)
	assert.Equal(t, "test.com", result.DomainName)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetBusinessUnitByID_NoUserID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.Background()
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	_, err := service.GetBusinessUnitByID(ctx, uuidStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetUserID)
}

func TestBusinessUnitService_GetBusinessUnitByID_InvalidUUID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")

	_, err := service.GetBusinessUnitByID(ctx, "invalid-uuid")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrInvalidUUIDFormat)
}

func TestBusinessUnitService_GetBusinessUnitByID_RepositoryError(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	mockRepo.On("GetBusinessUnitByID", ctx, mock.AnythingOfType("pgtype.UUID")).Return(repository.BusinessUnit{}, errors.New("database error"))

	_, err := service.GetBusinessUnitByID(ctx, uuidStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetBusinessUnit)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetAllDepartmentsInBusinessUnit_Success(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	validTime := time.Now()
	validUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}
	validTimePg := pgtype.Timestamptz{Time: validTime, Valid: true}

	mockDepartments := []repository.Department{
		{
			ID:             validUUID,
			BusinessUnitID: validUUID,
			Name:           "Test Department",
			IsActive:       pgtype.Bool{Bool: true, Valid: true},
			CreatedAt:      validTimePg,
			UpdatedAt:      validTimePg,
		},
	}

	mockRepo.On("GetAllDepartmentsInBusinessUnit", ctx, mock.AnythingOfType("pgtype.UUID")).Return(mockDepartments, nil)

	result, err := service.GetAllDepartmentsInBusinessUnit(ctx, uuidStr)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Department", result[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestBusinessUnitService_GetAllDepartmentsInBusinessUnit_NoUserID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.Background()
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	_, err := service.GetAllDepartmentsInBusinessUnit(ctx, uuidStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetUserID)
}

func TestBusinessUnitService_GetAllDepartmentsInBusinessUnit_InvalidUUID(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")

	_, err := service.GetAllDepartmentsInBusinessUnit(ctx, "invalid-uuid")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrInvalidUUIDFormat)
}

func TestBusinessUnitService_GetAllDepartmentsInBusinessUnit_RepositoryError(t *testing.T) {
	mockRepo := &MockQueries{}
	service := newBusinessUnitServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), utils.UserIDKey, "test-user")
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	mockRepo.On("GetAllDepartmentsInBusinessUnit", ctx, mock.AnythingOfType("pgtype.UUID")).Return([]repository.Department{}, errors.New("database error"))

	_, err := service.GetAllDepartmentsInBusinessUnit(ctx, uuidStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetDepartments)
	mockRepo.AssertExpectations(t)
}
