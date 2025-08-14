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
	ErrFailedToGetTenantID        = "failed to get tenant ID"
	ErrFailedToGetUserID          = "failed to get user ID"
	ErrFailedToGetBusinessUnits   = "failed to get business units from repository"
	ErrFailedToGetBusinessUnit    = "failed to get business unit from repository"
	ErrInvalidUUIDFormat          = "invalid UUID format"
	ErrFailedToCreateBusinessUnit = "failed to create business unit"

	// Department Service errors
	ErrFailedToGetDepartments        = "failed to get departments from repository"
	ErrFailedToGetDepartment         = "failed to get department from repository"
	ErrInvalidBusinessUnitUUIDFormat = "invalid business unit UUID format"
	ErrFailedToCreateDepartment      = "failed to create department"

	// User Service errors
	ErrFailedToGetUsers            = "failed to get users from repository"
	ErrFailedToGetUser             = "failed to get user from repository"
	ErrFailedToCreateUser          = "failed to create user in repository"
	ErrInvalidDepartmentUUIDFormat = "invalid department UUID format"
	ErrInvalidHomeTenantUUIDFormat = "invalid home tenant UUID format"
	ErrInvalidManagerUUIDFormat    = "invalid manager UUID format"
	ErrFailedToUpdateUser          = "failed to update user"

	// Permission Service errors
	ErrFailedToGetPermissions   = "failed to get permissions from repository"
	ErrFailedToGetPermission    = "failed to get permission from repository"
	ErrFailedToCreatePermission = "failed to create permission in repository"
	ErrFailedToUpdatePermission = "failed to update permission in repository"
	ErrFailedToDeletePermission = "failed to delete permission in repository"

	// Role Service errors
	ErrFailedToGetRoles   = "failed to get roles from repository"
	ErrFailedToGetRole    = "failed to get role from repository"
	ErrFailedToCreateRole = "failed to create role in repository"
	ErrFailedToUpdateRole = "failed to update role in repository"
	ErrFailedToDeleteRole = "failed to delete role in repository"

	// Scope Service errors
	ErrFailedToGetScopes   = "failed to get scopes from repository"
	ErrFailedToGetScope    = "failed to get scope from repository"
	ErrFailedToCreateScope = "failed to create scope in repository"
	ErrFailedToUpdateScope = "failed to update scope in repository"
	ErrFailedToDeleteScope = "failed to delete scope in repository"

	// RolePermission Service errors
	ErrFailedToGetRolePermissions   = "failed to get role permissions from repository"
	ErrFailedToGetRolePermission    = "failed to get role permission from repository"
	ErrFailedToCreateRolePermission = "failed to create role permission in repository"
	ErrFailedToUpdateRolePermission = "failed to update role permission in repository"
	ErrFailedToDeleteRolePermission = "failed to delete role permission in repository"
)

// Error variables
var (
	ErrEntraClientIDRequiredMsg     = fmt.Errorf("ENTRA_CLIENT_ID is required")
	ErrEntraClientSecretRequiredMsg = fmt.Errorf("ENTRA_CLIENT_SECRET is required")
	ErrEntraTenantIDRequiredMsg     = fmt.Errorf("ENTRA_TENANT_ID is required")
)

// Error messages
const (
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

	// Business Unit Controller error messages
	ErrBusinessUnitNotFoundMsg                = "Business unit not found"
	ErrBusinessUnitIDRequiredMsg              = "Business unit ID is required"
	ErrDomainNameRequiredMsg                  = "Domain name is required"
	ErrFailedToGetBusinessUnitByIDMsg         = "Failed to get business unit by ID"
	ErrFailedToGetBusinessUnitByDomainNameMsg = "Failed to get business unit by domain name"
	ErrFailedToRetrieveBusinessUnitsMsg       = "failed to retrieve business units"

	// Department Controller error messages
	ErrFailedToRetrieveDepartmentsMsg = "Failed to retrieve departments"
	ErrDepartmentNotFoundMsg          = "Department not found"
	ErrDepartmentIDRequiredMsg        = "Department ID is required"
	ErrDepartmentNameRequiredMsg      = "Department name is required"
	ErrFailedToGetDepartmentByIDMsg   = "Failed to get department by ID"
	ErrFailedToGetDepartmentByNameMsg = "Failed to get department by name"

	// User Controller error messages
	ErrFailedToRetrieveUsersMsg = "Failed to retrieve users"
	ErrUserNotFoundMsg          = "User not found"
	ErrInvalidRequestBodyMsg    = "Invalid request body"
	ErrFailedToCreateUserMsg    = "Failed to create user"

	// Permission Controller error messages
	ErrFailedToRetrievePermissionsMsg = "Failed to retrieve permissions"
	ErrPermissionNotFoundMsg          = "Permission not found"
	ErrPermissionIDRequiredMsg        = "Permission ID is required"
	ErrResourceRequiredMsg            = "Resource is required"
	ErrActionRequiredMsg              = "Action is required"
	ErrFailedToCreatePermissionMsg    = "Failed to create permission"
	ErrFailedToUpdatePermissionMsg    = "Failed to update permission"
	ErrFailedToDeletePermissionMsg    = "Failed to delete permission"
	ErrResourceNotFoundMsg            = "Resource not found"
	ErrNoPermissionsFoundMsg          = "No permissions found for this resource"
	ErrActionNotFoundMsg              = "Action not found"
	ErrResourceActionNotFoundMsg      = "Resource and action combination not found"

	// Role Controller error messages
	ErrFailedToRetrieveRolesMsg = "Failed to retrieve roles"
	ErrRoleNotFoundMsg          = "Role not found"
	ErrRoleIDRequiredMsg        = "Role ID is required"
	ErrFailedToCreateRoleMsg    = "Failed to create role"
	ErrFailedToUpdateRoleMsg    = "Failed to update role"
	ErrFailedToDeleteRoleMsg    = "Failed to delete role"

	// Scope Controller error messages
	ErrFailedToRetrieveScopesMsg = "Failed to retrieve scopes"
	ErrScopeNotFoundMsg          = "Scope not found"
	ErrScopeIDRequiredMsg        = "Scope ID is required"
	ErrFailedToCreateScopeMsg    = "Failed to create scope"
	ErrFailedToUpdateScopeMsg    = "Failed to update scope"
	ErrFailedToDeleteScopeMsg    = "Failed to delete scope"

	// RolePermission Controller error messages
	ErrFailedToRetrieveRolePermissionsMsg = "Failed to retrieve role permissions"
	ErrRolePermissionNotFoundMsg          = "Role permission not found"
	ErrRolePermissionIDRequiredMsg        = "Role permission ID is required"
	ErrFailedToCreateRolePermissionMsg    = "Failed to create role permission"
	ErrFailedToUpdateRolePermissionMsg    = "Failed to update role permission"
	ErrFailedToDeleteRolePermissionMsg    = "Failed to delete role permission"

	// RoleAssignment Service error messages
	ErrFailedToGetRoleAssignments   = "failed to get role assignments from repository"
	ErrFailedToGetRoleAssignment    = "failed to get role assignment from repository"
	ErrFailedToCreateRoleAssignment = "failed to create role assignment in repository"
	ErrFailedToUpdateRoleAssignment = "failed to update role assignment in repository"
	ErrFailedToDeleteRoleAssignment = "failed to delete role assignment in repository"
	ErrFailedToCheckUserPermission  = "failed to check user permission from repository"

	// RoleAssignment Controller error messages
	ErrFailedToRetrieveRoleAssignmentsMsg = "Failed to retrieve role assignments"
	ErrRoleAssignmentNotFoundMsg          = "Role assignment not found"
	ErrRoleAssignmentIDRequiredMsg        = "Role assignment ID is required"
	ErrUserIDRequiredMsg                  = "User ID is required"
	ErrFailedToCreateRoleAssignmentMsg    = "Failed to create role assignment"
	ErrFailedToUpdateRoleAssignmentMsg    = "Failed to update role assignment"
	ErrFailedToDeleteRoleAssignmentMsg    = "Failed to delete role assignment"
	ErrFailedToCheckPermissionMsg         = "Failed to check user permission"
	ErrFailedToGetUsersInDepartmentMsg    = "Failed to get users in department"
	ErrFailedToGetUserByIDMsg             = "Failed to get user by ID"
	ErrFailedToGetUserByEmailMsg          = "Failed to get user by email"
	ErrEmailIsRequiredMsg                 = "Email is required"

	// Form Category Controller error messages
	ErrFailedToGetFormCategories  = "failed to get form categories from repository"
	ErrFailedToGetFormCategory    = "failed to get form category from repository"
	ErrFailedToCreateFormCategory = "failed to create form category in repository"
	ErrFailedToUpdateFormCategory = "failed to update form category in repository"
	ErrFailedToDeleteFormCategory = "failed to delete form category in repository"
	ErrFormCategoryNotFound       = "form category not found"
	ErrFormCategoryNameRequired   = "form category name is required"
	ErrFormCategoryIDRequired     = "form category id is required"

	// Form Template errors
	ErrFailedToGetFormTemplates    = "Failed to get form templates"
	ErrFailedToGetFormTemplate     = "Failed to get form template"
	ErrFailedToCreateFormTemplate  = "Failed to create form template"
	ErrFailedToUpdateFormTemplate  = "Failed to update form template"
	ErrFailedToDeleteFormTemplate  = "Failed to delete form template"
	ErrFailedToPublishFormTemplate = "Failed to publish form template"
	ErrFormTemplateNotFound        = "Form template not found"
	ErrFormTemplateNameRequired    = "Form template name is required"
	ErrFormTemplateIDRequired      = "Form template ID is required"

	// Form Section errors
	ErrFailedToGetFormSections   = "Failed to get form sections"
	ErrFailedToGetFormSection    = "Failed to get form section"
	ErrFailedToCreateFormSection = "Failed to create form section"
	ErrFailedToUpdateFormSection = "Failed to update form section"
	ErrFailedToDeleteFormSection = "Failed to delete form section"
	ErrFormSectionNotFound       = "Form section not found"
	ErrFormSectionNameRequired   = "Form section name is required"
	ErrFormSectionIDRequired     = "Form section ID is required"

	// Request error messages
	ErrInvalidRequestBody = "invalid request body"
)
