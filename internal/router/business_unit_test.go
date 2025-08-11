package router

import (
	"testing"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessUnitController for testing
type MockBusinessUnitController struct {
	mock.Mock
}

func (m *MockBusinessUnitController) GetAllBusinessUnitsInTenant(c *gin.Context) {
	m.Called(c)
}

func (m *MockBusinessUnitController) GetBusinessUnitByDomainName(c *gin.Context) {
	m.Called(c)
}

func (m *MockBusinessUnitController) GetBusinessUnitByID(c *gin.Context) {
	m.Called(c)
}

func (m *MockBusinessUnitController) GetAllDepartmentsInBusinessUnit(c *gin.Context) {
	m.Called(c)
}

func TestNewBusinessUnitRouter(t *testing.T) {
	mockController := &controller.BusinessUnitController{}
	mockConfig := &config.Config{}

	router := NewBusinessUnitRouter(mockController, mockConfig)

	assert.NotNil(t, router)
	assert.Equal(t, mockController, router.controller)
	assert.Equal(t, mockConfig, router.config)
}

func TestBusinessUnitRouter_SetupBusinessUnitRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controller.BusinessUnitController{}
	mockConfig := &config.Config{
		OAuth: config.OAuthConfig{},
	}

	router := NewBusinessUnitRouter(mockController, mockConfig)

	ginEngine := gin.New()
	v1 := ginEngine.Group("/v1")

	// This should not panic
	assert.NotPanics(t, func() {
		router.SetupBusinessUnitRoutes(v1)
	})
}

func TestBusinessUnitRouter_RoutesRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controller.BusinessUnitController{}
	mockConfig := &config.Config{
		OAuth: config.OAuthConfig{},
	}

	router := NewBusinessUnitRouter(mockController, mockConfig)
	ginEngine := gin.New()
	v1 := ginEngine.Group("/v1")

	router.SetupBusinessUnitRoutes(v1)

	// Test that routes are registered
	routes := ginEngine.Routes()
	assert.True(t, len(routes) > 0)

	// Check specific routes exist
	routeExists := false
	for _, route := range routes {
		if route.Path == "/v1/business-units/" && route.Method == "GET" {
			routeExists = true
			break
		}
	}
	assert.True(t, routeExists)
}