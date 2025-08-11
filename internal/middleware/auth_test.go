package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/constants"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockJWKS for testing
type MockJWKS struct {
	mock.Mock
}

func (m *MockJWKS) Keyfunc(token *jwt.Token) (interface{}, error) {
	args := m.Called(token)
	return args.Get(0), args.Error(1)
}

func TestAuthMiddleWare_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create mock JWKS
	mockJWKS := &MockJWKS{}
	mockJWKS.On("Keyfunc", mock.Anything).Return([]byte("test-secret"), nil)

	// Create OAuth config
	oauthConfig := &config.OAuthConfig{
		ClientID:     "test-client-id",
		TenantID:     "test-tenant-id",
		JWKSEntra:    &keyfunc.JWKS{},
	}

	// Create test token with valid claims
	claims := &AzureADClaims{
		Audience: "test-client-id",
		Issuer:   "https://login.microsoftonline.com/test-tenant-id/v2.0",
		TID:      "test-tenant-id",
		OID:      "test-user-id",
		Name:     "Test User",
		Email:    "test@example.com",
		Expiry:   time.Now().Add(time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("test-secret"))

	// Create test request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	// Test middleware
	middleware := AuthMiddleWare(oauthConfig)
	middleware(c)

	// Assertions
	assert.False(t, c.IsAborted())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleWare_MissingOAuthConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)

	middleware := AuthMiddleWare(nil)
	middleware(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAuthMiddleWare_MissingAuthHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	oauthConfig := &config.OAuthConfig{
		ClientID:  "test-client-id",
		TenantID:  "test-tenant-id",
		JWKSEntra: &keyfunc.JWKS{},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)

	middleware := AuthMiddleWare(oauthConfig)
	middleware(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestValidateAuthHeader_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer valid-token")

	token, err := validateAuthHeader(c)

	assert.NoError(t, err)
	assert.Equal(t, "valid-token", token)
}

func TestValidateAuthHeader_MissingHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/test", nil)

	_, err := validateAuthHeader(c)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrMissingAuthHeaderMsg)
}

func TestValidateAuthHeader_InvalidFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "InvalidFormat")

	_, err := validateAuthHeader(c)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrMissingAuthHeaderMsg)
}

func TestValidateAudience_StringAudience(t *testing.T) {
	result := validateAudience("test-audience", "test-audience")
	assert.True(t, result)

	result = validateAudience("wrong-audience", "test-audience")
	assert.False(t, result)
}

func TestValidateAudience_SliceAudience(t *testing.T) {
	audiences := []interface{}{"aud1", "test-audience", "aud3"}
	result := validateAudience(audiences, "test-audience")
	assert.True(t, result)

	result = validateAudience(audiences, "wrong-audience")
	assert.False(t, result)
}

func TestValidateTokenExpiry_ValidToken(t *testing.T) {
	claims := &AzureADClaims{
		Expiry: time.Now().Add(time.Hour).Unix(),
	}

	err := validateTokenExpiry(claims)
	assert.NoError(t, err)
}

func TestValidateTokenExpiry_ExpiredToken(t *testing.T) {
	claims := &AzureADClaims{
		Expiry: time.Now().Add(-time.Hour).Unix(),
	}

	err := validateTokenExpiry(claims)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrTokenExpiredMsg)
}

func TestValidateTokenExpiry_NoExpiry(t *testing.T) {
	claims := &AzureADClaims{}

	err := validateTokenExpiry(claims)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), constants.ErrTokenExpiryNotSetMsg)
}