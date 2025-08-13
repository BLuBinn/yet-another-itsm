package constants

// Success messages
const (
	SuccessMsg = "success"

	// Health Controller success messages
	SuccessMsgHealthCheck = "Health check completed"

	// Business Unit Controller success messages
	SuccessMsgGetAllBusinessUnits         = "Successfully retrieved all business units"
	SuccessMsgGetBusinessUnitByDomainName = "Successfully retrieved business unit by domain name"
	SuccessMsgGetBusinessUnitByID         = "Successfully retrieved business unit by ID"

	// Department Controller success messages
	SuccessMsgGetAllDepartments   = "Successfully retrieved all departments"
	SuccessMsgGetDepartmentByID   = "Successfully retrieved department by ID"
	SuccessMsgGetDepartmentByName = "Successfully retrieved department by name"

	// User Controller success messages
	SuccessMsgGetCurrentUser          = "Successfully retrieved current user"
	SuccessMsgGetAllUsersInDepartment = "Successfully retrieved all users in department"
	SuccessMsgGetUserByID             = "Successfully retrieved user by ID"
	SuccessMsgGetUserByEmail          = "Successfully retrieved user by email"
	SuccessMsgCreateUser              = "Successfully created user"

	// Form Template Controller success messages
	SuccessGetFormTemplates   = "Successfully retrieved all form templates"
	SuccessCreateFormTemplate = "Successfully created form template"
	SuccessUpdateFormTemplate = "Successfully updated form template"
	SuccessDeleteFormTemplate = "Successfully deleted form template"

	// Form Section Controller success messages
	SuccessGetFormSections   = "Successfully retrieved all form sections"
	SuccessCreateFormSection = "Successfully created form section"
	SuccessUpdateFormSection = "Successfully updated form section"
	SuccessDeleteFormSection = "Successfully deleted form section"
)
