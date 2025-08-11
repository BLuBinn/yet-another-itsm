package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessUnitService for testing
type MockBusinessUnitService struct {
	mock.Mock
}

func (m *MockBusinessUnitService) GetAllBusinessUnitsInTenant(ctx context.Context) ([]*dtos.BusinessUnit, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*dtos.BusinessUnit), args.Error(1)
}

func (m *MockBusinessUnitService) GetBusinessUnitByDomainName(ctx context.Context, domainName string) (*dtos.BusinessUnit, error) {
	args := m.Called(ctx, domainName)
	return args.Get(0).(*dtos.BusinessUnit), args.Error(1)
}

func (m *MockBusinessUnitService) GetBusinessUnitByID(ctx context.Context, id string) (*dtos.BusinessUnit, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dtos.BusinessUnit), args.Error(1)
}

func (m *MockBusinessUnitService) GetAllDepartmentsInBusinessUnit(ctx context.Context, businessUnitID string) ([]*dtos.Department, error) {
	args := m.Called(ctx, businessUnitID)
	return args.Get(0).([]*dtos.Department), args.Error(1)
}

func TestNewBusinessUnitController(t *testing.T) {
	mockServices := &service.Services{}
	controller := NewBusinessUnitController(mockServices)

	assert.NotNil(t, controller)
	assert.Equal(t, mockServices, controller.services)
}

func TestBusinessUnitController_GetAllBusinessUnitsInTenant_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockBusinessUnitService{}
	mockServices := &service.Services{
		BusinessUnit: mockService,
	}

	controller := NewBusinessUnitController(mockServices)

	// Mock data
	mockBusinessUnits := []*dtos.BusinessUnit{
		{
			DomainName: "test.com",
			TenantID:   "tenant-1",
			Name:       "Test Business Unit",
		},
	}

	mockService.On("GetAllBusinessUnitsInTenant", mock.Anything).Return(mockBusinessUnits, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units", nil)

	controller.GetAllBusinessUnitsInTenant(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestBusinessUnitController_GetAllBusinessUnitsInTenant_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockBusinessUnitService{}
	mockServices := &service.Services{
		BusinessUnit: mockService,
	}

	controller := NewBusinessUnitController(mockServices)

	mockService.On("GetAllBusinessUnitsInTenant", mock.Anything).Return([]*dtos.BusinessUnit{}, errors.New("database error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units", nil)

	controller.GetAllBusinessUnitsInTenant(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestBusinessUnitController_GetBusinessUnitByDomainName_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockBusinessUnitService{}
	mockServices := &service.Services{
		BusinessUnit: mockService,
	}

	controller := NewBusinessUnitController(mockServices)

	mockBusinessUnit := &dtos.BusinessUnit{
		DomainName: "test.com",
		TenantID:   "tenant-1",
		Name:       "Test Business Unit",
	}

	mockService.On("GetBusinessUnitByDomainName", mock.Anything, "test.com").Return(mockBusinessUnit, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units?domain=test.com", nil)

	controller.GetBusinessUnitByDomainName(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestBusinessUnitController_GetBusinessUnitByDomainName_MissingDomain(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockServices := &service.Services{}
	controller := NewBusinessUnitController(mockServices)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units", nil)

	controller.GetBusinessUnitByDomainName(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBusinessUnitController_GetBusinessUnitByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockBusinessUnitService{}
	mockServices := &service.Services{
		BusinessUnit: mockService,
	}

	controller := NewBusinessUnitController(mockServices)

	mockBusinessUnit := &dtos.BusinessUnit{
		DomainName: "test.com",
		TenantID:   "tenant-1",
		Name:       "Test Business Unit",
	}

	mockService.On("GetBusinessUnitByID", mock.Anything, "123").Return(mockBusinessUnit, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units/123", nil)
	c.Params = gin.Params{{Key: "businessUnitId", Value: "123"}}

	controller.GetBusinessUnitByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestBusinessUnitController_GetAllDepartmentsInBusinessUnit_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockBusinessUnitService{}
	mockServices := &service.Services{
		BusinessUnit: mockService,
	}

	controller := NewBusinessUnitController(mockServices)

	mockDepartments := []*dtos.Department{
		{
			BusinessUnitID: "123",
			Name:           "Test Department",
		},
	}

	mockService.On("GetAllDepartmentsInBusinessUnit", mock.Anything, "123").Return(mockDepartments, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/v1/business-units/123/departments", nil)
	c.Params = gin.Params{{Key: "businessUnitId", Value: "123"}}

	controller.GetAllDepartmentsInBusinessUnit(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}