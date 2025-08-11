package service

import (
	"context"
	"errors"
	"testing"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuerier implements the full Querier interface for testing
type MockQuerier struct {
	mock.Mock
}

// Implement all Querier interface methods
func (m *MockQuerier) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repository.User), args.Error(1)
}

func (m *MockQuerier) GetAllBusinessUnitsInTenant(ctx context.Context, tenantID string) ([]repository.BusinessUnit, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0).([]repository.BusinessUnit), args.Error(1)
}

func (m *MockQuerier) GetAllDepartmentsInBusinessUnit(ctx context.Context, businessUnitID pgtype.UUID) ([]repository.Department, error) {
	args := m.Called(ctx, businessUnitID)
	return args.Get(0).([]repository.Department), args.Error(1)
}

func (m *MockQuerier) GetAllUsersInDepartment(ctx context.Context, departmentID pgtype.UUID) ([]repository.User, error) {
	args := m.Called(ctx, departmentID)
	return args.Get(0).([]repository.User), args.Error(1)
}

func (m *MockQuerier) GetBusinessUnitByDomainName(ctx context.Context, domainName string) (repository.BusinessUnit, error) {
	args := m.Called(ctx, domainName)
	return args.Get(0).(repository.BusinessUnit), args.Error(1)
}

func (m *MockQuerier) GetBusinessUnitByID(ctx context.Context, id pgtype.UUID) (repository.BusinessUnit, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.BusinessUnit), args.Error(1)
}

func (m *MockQuerier) GetDepartmentByID(ctx context.Context, id pgtype.UUID) (repository.Department, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.Department), args.Error(1)
}

func (m *MockQuerier) GetDepartmentByName(ctx context.Context, params repository.GetDepartmentByNameParams) (repository.Department, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(repository.Department), args.Error(1)
}

func (m *MockQuerier) GetUserByEmail(ctx context.Context, mail string) (repository.User, error) {
	args := m.Called(ctx, mail)
	return args.Get(0).(repository.User), args.Error(1)
}

func (m *MockQuerier) GetUserByID(ctx context.Context, id pgtype.UUID) (repository.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repository.User), args.Error(1)
}

func TestNewDepartmentService(t *testing.T) {
	mockRepo := &MockQuerier{}
	// Create a repository.Queries with the mock
	// You'll need to modify this based on how repository.Queries is structured
	// For now, let's create a wrapper function
	service := newDepartmentServiceWithMock(mockRepo)

	assert.NotNil(t, service)
	assert.IsType(t, &departmentService{}, service)
}

// Helper function to create service with mock
func newDepartmentServiceWithMock(querier repository.Querier) DepartmentService {
	return &departmentService{
		repo: &repository.Queries{}, // This needs to be adapted based on actual structure
	}
}

func TestDepartmentService_GetDepartmentByID_Success(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), "user_id", "test-user")
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"

	mockDepartment := repository.Department{
		ID:             pgtype.UUID{},
		BusinessUnitID: pgtype.UUID{},
		Name:           "Test Department",
		IsActive:       pgtype.Bool{Bool: true, Valid: true},
	}

	mockRepo.On("GetDepartmentByID", mock.Anything, mock.AnythingOfType("pgtype.UUID")).Return(mockDepartment, nil)

	dept, err := service.GetDepartmentByID(ctx, uuidStr)

	assert.NoError(t, err)
	assert.NotNil(t, dept)
	assert.Equal(t, "Test Department", dept.Name)
	mockRepo.AssertExpectations(t)
}

func TestDepartmentService_GetDepartmentByID_NoUserID(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.Background()

	_, err := service.GetDepartmentByID(ctx, "550e8400-e29b-41d4-a716-446655440000")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetUserID)
}

func TestDepartmentService_GetDepartmentByID_InvalidUUID(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), "user_id", "test-user")

	_, err := service.GetDepartmentByID(ctx, "invalid-uuid")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrInvalidUUIDFormat)
}

func TestDepartmentService_GetDepartmentByName_Success(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), "user_id", "test-user")
	businessUnitID := "550e8400-e29b-41d4-a716-446655440000"
	departmentName := "Test Department"

	mockDepartment := repository.Department{
		ID:             pgtype.UUID{},
		BusinessUnitID: pgtype.UUID{},
		Name:           departmentName,
		IsActive:       pgtype.Bool{Bool: true, Valid: true},
	}

	mockRepo.On("GetDepartmentByName", ctx, mock.AnythingOfType("repository.GetDepartmentByNameParams")).Return(mockDepartment, nil)

	result, err := service.GetDepartmentByName(ctx, departmentName, businessUnitID)

	assert.NoError(t, err)
	assert.Equal(t, departmentName, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestDepartmentService_GetDepartmentByName_InvalidBusinessUnitUUID(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), "user_id", "test-user")

	_, err := service.GetDepartmentByName(ctx, "Test Department", "invalid-uuid")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrInvalidBusinessUnitUUIDFormat)
}

func TestDepartmentService_GetDepartmentByName_RepositoryError(t *testing.T) {
	mockRepo := &MockQuerier{}
	service := newDepartmentServiceWithMock(mockRepo)

	ctx := context.WithValue(context.Background(), "user_id", "test-user")
	businessUnitID := "550e8400-e29b-41d4-a716-446655440000"
	departmentName := "Test Department"

	mockRepo.On("GetDepartmentByName", ctx, mock.AnythingOfType("repository.GetDepartmentByNameParams")).Return(repository.Department{}, errors.New("database error"))

	_, err := service.GetDepartmentByName(ctx, departmentName, businessUnitID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrFailedToGetDepartment)
	mockRepo.AssertExpectations(t)
}