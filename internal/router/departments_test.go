package router

import (
	"testing"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewDepartmentRouter(t *testing.T) {
	mockController := &controller.DepartmentController{}
	mockConfig := &config.Config{}

	router := NewDepartmentRouter(mockController, mockConfig)

	assert.NotNil(t, router)
	assert.Equal(t, mockController, router.controller)
	assert.Equal(t, mockConfig, router.config)
}

func TestDepartmentRouter_SetupDepartmentRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controller.DepartmentController{}
	mockConfig := &config.Config{
		OAuth: config.OAuthConfig{},
	}

	router := NewDepartmentRouter(mockController, mockConfig)

	ginEngine := gin.New()
	v1 := ginEngine.Group("/v1")

	// This should not panic
	assert.NotPanics(t, func() {
		router.SetupDepartmentRoutes(v1)
	})
}

func TestDepartmentRouter_RoutesRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controller.DepartmentController{}
	mockConfig := &config.Config{
		OAuth: config.OAuthConfig{},
	}

	router := NewDepartmentRouter(mockController, mockConfig)
	ginEngine := gin.New()
	v1 := ginEngine.Group("/v1")

	router.SetupDepartmentRoutes(v1)

	// Test that routes are registered
	routes := ginEngine.Routes()
	assert.True(t, len(routes) > 0)

	// Check specific routes exist
	routeExists := false
	for _, route := range routes {
		if route.Path == "/v1/departments/" && route.Method == "GET" {
			routeExists = true
			break
		}
	}
	assert.True(t, routeExists)
}