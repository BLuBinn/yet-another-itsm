package constants

import "fmt"

// Error codes
const (
	// Error codes for common errors
	ErrCodeInternal         = "INTERNAL_ERROR"
	ErrCodeNotFound         = "NOT_FOUND"
	ErrCodeBadRequest       = "BAD_REQUEST"
	ErrCodeUnauthorized     = "UNAUTHORIZED"
	ErrCodeForbidden        = "FORBIDDEN"
	ErrCodeConflict         = "CONFLICT"
	ErrCodeValidation       = "VALIDATION_ERROR"
	ErrCodeTimeout          = "TIMEOUT"
	ErrCodeUnavailable      = "SERVICE_UNAVAILABLE"
	ErrCodeMethodNotAllowed = "METHOD_NOT_ALLOWED"

	// Common errors
	ErrTenantIDNotFound    = "tenant ID not found in context"
	ErrUserIDNotFound      = "user ID not found in context"
	ErrUserNameNotFound    = "user name not found in context"
	ErrAccessTokenNotFound = "access token not found in context"

	// Health Service errors
	ErrDatabaseHealthCheckFailed = "Database health check failed"

	HealthStatusOK         = "ok"
	HealthStatusDegraded   = "degraded"
	HealthStatusUnhealthy  = "unhealthy"
	HealthServiceDatabase  = "database"
	HealthServiceHealthy   = "healthy"
	HealthServiceUnhealthy = "unhealthy"

	// Graph Service errors
	ErrCouldNotCreateOBOCredential = "could not create OBO credential"
	ErrCouldNotCreateGraphClient   = "could not create Graph client"
	ErrFailedToGetCurrentUser      = "failed to get current user"

	// Business Unit Service errors
	ErrFailedToGetTenantID      = "failed to get tenant ID"
	ErrFailedToGetUserID        = "failed to get user ID"
	ErrFailedToGetBusinessUnits = "failed to get business units from repository"
	ErrFailedToGetBusinessUnit  = "failed to get business unit from repository"
	ErrInvalidUUIDFormat        = "invalid UUID format"

	// Department Service errors
	ErrFailedToGetDepartments        = "failed to get departments from repository"
	ErrFailedToGetDepartment         = "failed to get department from repository"
	ErrInvalidBusinessUnitUUIDFormat = "invalid business unit UUID format"

	// User Service errors
	ErrFailedToGetUsers            = "failed to get users from repository"
	ErrFailedToGetUser             = "failed to get user from repository"
	ErrFailedToCreateUser          = "failed to create user in repository"
	ErrInvalidDepartmentUUIDFormat = "invalid department UUID format"
	ErrInvalidHomeTenantUUIDFormat = "invalid home tenant UUID format"
	ErrInvalidManagerUUIDFormat    = "invalid manager UUID format"
)

// Error variables
var (
	ErrEntraClientIDRequiredMsg     = fmt.Errorf("ENTRA_CLIENT_ID is required")
	ErrEntraClientSecretRequiredMsg = fmt.Errorf("ENTRA_CLIENT_SECRET is required")
	ErrEntraTenantIDRequiredMsg     = fmt.Errorf("ENTRA_TENANT_ID is required")
)

// Error messages
const (
	// Common error messages
	ErrInternalServerMsg   = "Internal server error"
	ErrBadRequestMsg       = "Bad request"
	ErrUnauthorizedMsg     = "Unauthorized"
	ErrForbiddenMsg        = "Forbidden"
	ErrNotFoundMsg         = "Not found"
	ErrConflictMsg         = "Conflict"
	ErrValidationMsg       = "Validation error"
	ErrTimeoutMsg          = "Timeout"
	ErrUnavailableMsg      = "Service unavailable"
	ErrMethodNotAllowedMsg = "Method not allowed"
	ErrRouteNotFoundMsg    = "Route not found"

	// Config OAuth error messages
	ErrFailedToLoadJWKSMsg       = "failed to load JWKS for Entra ID: %w"
	ErrJWKSRefreshErrorMsg       = "[Entra JWKS] Error refreshing JWKS: %v"
	ErrOAuthConfigInitializedMsg = "OAuth config initialized successfully"

	// Logger error messages
	ErrFailedToInitializeOAuthMsg = "failed to initialize OAuth: %w"
	ErrInvalidLogLevelMsg         = "invalid log level, using info"
	ErrNoEnvFileFoundMsg          = "no .env file found or error loading .env file, using system environment variables"

	// Middleware error messages
	ErrAuthServiceNotAvailableMsg       = "Authentication service not available"
	ErrMissingAuthHeaderMsg             = "Missing or invalid Authorization header"
	ErrInvalidTokenMsg                  = "Invalid token"
	ErrInvalidTokenClaimsMsg            = "Invalid token claims"
	ErrFailedToParseJWTTokenMsg         = "Failed to parse JWT token"
	ErrTokenExpiredMsg                  = "Token expired"
	ErrInvalidTokenIssuerMsg            = "Invalid token issuer"
	ErrInvalidTokenAudienceMsg          = "Invalid token audience"
	ErrInvalidTenantIDMsg               = "Invalid tenant ID"
	ErrUserAuthenticatedSuccessfullyMsg = "User authenticated successfully"
	ErrTokenExpiryNotSetMsg             = "token expiry not set"
	ErrTokenExpiredAtMsg                = "token expired at %v, current time %v"

	// Graph Controller error messages
	ErrAccessTokenNotFoundMsg        = "Access token not found"
	ErrFailedToGetUserInformationMsg = "Failed to get user information"

	// Health Controller error messages
	ErrHealthCheckFailedMsg      = "Health check failed"
	ErrSystemHealthyMsg          = "System is healthy"
	ErrSystemPartiallyHealthyMsg = "System is partially healthy"
	ErrSystemUnhealthyMsg        = "System is unhealthy"

	// Graph Controller error messages
	ErrFailedToRetrieveBusinessUnitsMsg = "Failed to retrieve business units"

	// Business Unit Controller error messages
	ErrBusinessUnitNotFoundMsg                = "Business unit not found"
	ErrBusinessUnitIDRequiredMsg              = "Business unit ID is required"
	ErrDomainNameRequiredMsg                  = "Domain name is required"
	ErrFailedToGetBusinessUnitByIDMsg         = "Failed to get business unit by ID"
	ErrFailedToGetBusinessUnitByDomainNameMsg = "Failed to get business unit by domain name"

	// Department Controller error messages
	ErrFailedToRetrieveDepartmentsMsg = "Failed to retrieve departments"
	ErrDepartmentNotFoundMsg          = "Department not found"
	ErrDepartmentIDRequiredMsg        = "Department ID is required"
	ErrDepartmentNameRequiredMsg      = "Department name is required"
	ErrFailedToGetDepartmentByIDMsg   = "Failed to get department by ID"
	ErrFailedToGetDepartmentByNameMsg = "Failed to get department by name"

	// User Controller error messages
	ErrFailedToRetrieveUsersMsg        = "Failed to retrieve users"
	ErrFailedToGetUsersInDepartmentMsg = "Failed to get users in department"
	ErrFailedToGetUserByIDMsg          = "Failed to get user by ID"
	ErrFailedToGetUserByEmailMsg       = "Failed to get user by email"
	ErrUserNotFoundMsg                 = "User not found"
	ErrInvalidRequestBodyMsg           = "Invalid request body"
	ErrFailedToCreateUserMsg           = "Failed to create user"
	ErrEmailIsRequiredMsg              = "Email is required"
)
