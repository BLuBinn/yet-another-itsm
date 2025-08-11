package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad_Success(t *testing.T) {
	// Set required environment variables
	os.Setenv("ENTRA_CLIENT_ID", "test-client-id")
	os.Setenv("ENTRA_CLIENT_SECRET", "test-client-secret")
	os.Setenv("ENTRA_TENANT_ID", "test-tenant-id")
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("LOG_FORMAT", "json")

	defer func() {
		os.Unsetenv("ENTRA_CLIENT_ID")
		os.Unsetenv("ENTRA_CLIENT_SECRET")
		os.Unsetenv("ENTRA_TENANT_ID")
		os.Unsetenv("PORT")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("LOG_LEVEL")
		os.Unsetenv("LOG_FORMAT")
	}()

	cfg, err := Load()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, "5432", cfg.Database.Port)
	assert.Equal(t, "postgres", cfg.Database.User)
	assert.Equal(t, "password", cfg.Database.Password)
	assert.Equal(t, "testdb", cfg.Database.DBName)
	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "json", cfg.Logger.Format)
	assert.Equal(t, "test-client-id", cfg.OAuth.ClientID)
	assert.Equal(t, "test-client-secret", cfg.OAuth.ClientSecret)
	assert.Equal(t, "test-tenant-id", cfg.OAuth.TenantID)
}

func TestLoad_MissingClientID(t *testing.T) {
	os.Unsetenv("ENTRA_CLIENT_ID")
	os.Setenv("ENTRA_CLIENT_SECRET", "test-client-secret")
	os.Setenv("ENTRA_TENANT_ID", "test-tenant-id")

	defer func() {
		os.Unsetenv("ENTRA_CLIENT_SECRET")
		os.Unsetenv("ENTRA_TENANT_ID")
	}()

	_, err := Load()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ENTRA_CLIENT_ID is required")
}

func TestLoad_DefaultValues(t *testing.T) {
	// Set only required OAuth variables
	os.Setenv("ENTRA_CLIENT_ID", "test-client-id")
	os.Setenv("ENTRA_CLIENT_SECRET", "test-client-secret")
	os.Setenv("ENTRA_TENANT_ID", "test-tenant-id")

	defer func() {
		os.Unsetenv("ENTRA_CLIENT_ID")
		os.Unsetenv("ENTRA_CLIENT_SECRET")
		os.Unsetenv("ENTRA_TENANT_ID")
	}()

	cfg, err := Load()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	// Check default values
	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, 10*time.Second, cfg.Server.ReadTimeout)
	assert.Equal(t, 10*time.Second, cfg.Server.WriteTimeout)
	assert.Equal(t, 120*time.Second, cfg.Server.IdleTimeout)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, "5432", cfg.Database.Port)
	assert.Equal(t, "postgres", cfg.Database.User)
	assert.Equal(t, "postgres", cfg.Database.Password)
	assert.Equal(t, "msn_map_api", cfg.Database.DBName)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
	assert.Equal(t, int32(25), cfg.Database.MaxConns)
	assert.Equal(t, int32(5), cfg.Database.MinConns)
	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "json", cfg.Logger.Format)
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	result := getEnv("TEST_VAR", "default_value")
	assert.Equal(t, "test_value", result)

	result = getEnv("NON_EXISTENT_VAR", "default_value")
	assert.Equal(t, "default_value", result)
}

func TestGetIntEnv(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	defer os.Unsetenv("TEST_INT")

	result := getIntEnv("TEST_INT", 10)
	assert.Equal(t, 42, result)

	result = getIntEnv("NON_EXISTENT_INT", 10)
	assert.Equal(t, 10, result)

	// Test invalid integer
	os.Setenv("INVALID_INT", "not_a_number")
	defer os.Unsetenv("INVALID_INT")

	result = getIntEnv("INVALID_INT", 10)
	assert.Equal(t, 10, result)
}

func TestGetDurationEnv(t *testing.T) {
	os.Setenv("TEST_DURATION", "30s")
	defer os.Unsetenv("TEST_DURATION")

	result := getDurationEnv("TEST_DURATION", 10*time.Second)
	assert.Equal(t, 30*time.Second, result)

	result = getDurationEnv("NON_EXISTENT_DURATION", 10*time.Second)
	assert.Equal(t, 10*time.Second, result)

	// Test invalid duration
	os.Setenv("INVALID_DURATION", "not_a_duration")
	defer os.Unsetenv("INVALID_DURATION")

	result = getDurationEnv("INVALID_DURATION", 10*time.Second)
	assert.Equal(t, 10*time.Second, result)
}