package controller

import (
	"context"
	"net/http"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/model"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/rs/zerolog/log"

	responseModel "yet-another-itsm/internal/dtos"
)

type UserController struct {
	userService service.UserService
	services    *service.Services
}

func NewUserController(services *service.Services) *UserController {
	return &UserController{
		userService: services.User,
		services:    services,
	}
}

// GetAllUsersInDepartment godoc
// @Summary Get all users in department
// @Description Get all users in a specific department
// @Tags users
// @Accept json
// @Produce json
// @Param departmentId path string true "Department ID"
// @Success 200 {object} dtos.UsersListResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/departments/{departmentId}/users [get]
func (uc *UserController) GetAllUsersInDepartment(c *gin.Context) {
	departmentID := c.Param("departmentId")

	log.Info().
		Str("controller", "UserController").
		Str("endpoint", "GetAllUsersInDepartment").
		Str("method", c.Request.Method).
		Str("department_id", departmentID).
		Msg("Getting all users in department")

	ctx := c.Request.Context()

	users, err := uc.userService.GetAllUsersInDepartment(ctx, departmentID)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToGetUsersInDepartmentMsg)
		utils.SendNotFound(c, constants.ErrDepartmentNotFoundMsg)
		return
	}

	var userResponses []responseModel.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, *user.ToResponse())
	}

	response := responseModel.NewUsersListResponse(userResponses, 1, len(userResponses), int64(len(userResponses))) // Thay vì dtos.NewUsersListResponse

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetAllUsersInDepartment, response)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a specific user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dtos.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/users/{userId} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("userId")

	log.Info().
		Str("controller", "UserController").
		Str("endpoint", "GetUserByID").
		Str("method", c.Request.Method).
		Str("id", id).
		Msg("Getting user by ID")

	ctx := c.Request.Context()

	user, err := uc.userService.GetUserByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToGetUserByIDMsg)
		utils.SendNotFound(c, constants.ErrUserNotFoundMsg)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetUserByID, user.ToResponse())
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Get a specific user by email
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} dtos.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/users/email [get]
func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		utils.SendBadRequest(c, constants.ErrEmailIsRequiredMsg)
		return
	}

	log.Info().
		Str("controller", "UserController").
		Str("endpoint", "GetUserByEmail").
		Str("method", c.Request.Method).
		Str("email", email).
		Msg("Getting user by email")

	ctx := c.Request.Context()

	user, err := uc.userService.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToGetUserByEmailMsg)
		utils.SendNotFound(c, constants.ErrUserNotFoundMsg)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetUserByEmail, user.ToResponse())
}

// GetCurrentUser godoc
// @Summary Get user information
// @Description Get user information and create user if not exists
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserInfoResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /v1/users/me [get]
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	log.Info().
		Str("controller", "UserController").
		Str("endpoint", "GetCurrentUser").
		Str("method", c.Request.Method).
		Msg("Getting current user from user")

	accessToken, err := utils.GetAccessToken(c.Request.Context())
	if err != nil {
		utils.SendUnauthorized(c, constants.ErrAccessTokenNotFoundMsg)
		return
	}

	user, err := uc.services.Graph.GetCurrentUser(accessToken)
	if err != nil {
		utils.SendUnauthorized(c, constants.ErrFailedToGetUserInformationMsg)
		return
	}

	ctx := c.Request.Context()
	userEmail := utils.GetStringValue(user.GetMail())

	// Check if user exists and get business unit info using GetOrCreate
	existingUser, businessUnitName := uc.handleExistingUserWithGetOrCreate(ctx, userEmail)

	// Update last login asynchronously
	uc.updateLastLoginAsync(userEmail)

	// Create user if not exists using GetOrCreate methods
	if existingUser == nil {
		if err := uc.createNewUserWithGetOrCreate(ctx, c, user, userEmail); err != nil {
			return // Error already handled in createNewUser
		}
	}

	// Build and send response
	userInfoResponse := uc.buildUserInfoResponse(user, businessUnitName)
	log.Info().
		Str("user_id", userInfoResponse.ID).
		Msg(constants.SuccessMsgGetCurrentUser)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetCurrentUser, userInfoResponse)
}

// handleExistingUserWithGetOrCreate checks if user exists and gets business unit name using GetOrCreate
func (uc *UserController) handleExistingUserWithGetOrCreate(ctx context.Context, userEmail string) (*responseModel.User, string) {
	existingUser, err := uc.services.User.GetUserByEmail(ctx, userEmail)
	businessUnitName := uc.getBusinessUnitNameWithGetOrCreate(ctx, userEmail)

	if err != nil {
		return nil, businessUnitName
	}
	return existingUser, businessUnitName
}

// getBusinessUnitNameWithGetOrCreate extracts domain from email and gets business unit name using GetOrCreate
func (uc *UserController) getBusinessUnitNameWithGetOrCreate(ctx context.Context, userEmail string) string {
	domainName, isValid := utils.ExtractDomainFromEmail(userEmail)
	if !isValid {
		return ""
	}

	// Sử dụng GetOrCreateBusinessUnitByDomainName thay vì GetBusinessUnitByDomainName
	businessUnit, err := uc.services.BusinessUnit.GetOrCreateBusinessUnitByDomainName(ctx, domainName)
	if err != nil {
		log.Error().Err(err).
			Str("domain_name", domainName).
			Str("user_email", userEmail).
			Msg("Failed to get or create business unit by domain")
		return ""
	}

	log.Info().
		Str("domain_name", domainName).
		Str("business_unit_name", businessUnit.Name).
		Str("user_email", userEmail).
		Msg("Successfully got or created business unit by domain")

	return businessUnit.Name
}

// updateLastLoginAsync updates user's last login time in background
func (uc *UserController) updateLastLoginAsync(userEmail string) {
	go func() {
		if updateErr := uc.services.User.UpdateUserLastLogin(context.Background(), userEmail); updateErr != nil {
			log.Error().Err(updateErr).
				Str("email", userEmail).
				Msg("Failed to update last login in background")
		}
	}()
}

// createNewUserWithGetOrCreate creates a new user using GetOrCreate methods for department and business unit
func (uc *UserController) createNewUserWithGetOrCreate(ctx context.Context, c *gin.Context, user models.Userable, userEmail string) error {
	log.Info().
		Str("email", userEmail).
		Msg("User not found in database, creating new user")

	// Sử dụng GetOrCreate methods thay vì logic phức tạp
	departmentID := uc.getDepartmentIDWithGetOrCreate(ctx, user)
	businessUnitID := uc.getBusinessUnitIDWithGetOrCreate(ctx, userEmail)

	tenantID, err := utils.GetTenantID(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant ID from context")
		utils.SendInternalServerError(c, constants.ErrFailedToGetTenantID)
		return err
	}

	createUserReq := &responseModel.CreateUserRequest{
		AzureAdObjectID: utils.GetStringValue(user.GetId()),
		HomeTenantID:    tenantID,
		Mail:            userEmail,
		DisplayName:     utils.GetStringValue(user.GetDisplayName()),
		GivenName:       utils.GetStringValue(user.GetGivenName()),
		SurName:         utils.GetStringValue(user.GetSurname()),
		JobTitle:        utils.GetStringValue(user.GetJobTitle()),
		OfficeLocation:  utils.GetStringValue(user.GetOfficeLocation()),
		DepartmentID:    departmentID,
		BusinessUnitID:  businessUnitID,
		ManagerID:       responseModel.GetManagerId(user),
		Status:          model.UserStatusActive,
	}

	_, createErr := uc.services.User.CreateUser(ctx, createUserReq)
	if createErr != nil {
		log.Error().Err(createErr).Msg("Failed to create user in database")
	} else {
		log.Info().
			Str("email", userEmail).
			Msg("Successfully created user in database")
	}

	return nil
}

// getDepartmentIDWithGetOrCreate gets department ID using GetOrCreateDepartmentByName
func (uc *UserController) getDepartmentIDWithGetOrCreate(ctx context.Context, user models.Userable) string {
	departmentName := utils.GetStringValue(user.GetDepartment())
	if departmentName == "" {
		return ""
	}

	// Sử dụng GetOrCreateDepartmentByName thay vì logic phức tạp
	department, err := uc.services.Department.GetOrCreateDepartmentByName(ctx, departmentName)
	if err != nil {
		log.Error().Err(err).
			Str("department_name", departmentName).
			Msg("Failed to get or create department")
		return ""
	}

	log.Info().
		Str("department_name", departmentName).
		Str("department_id", department.ID).
		Msg("Successfully got or created department")

	return department.ID
}

// getBusinessUnitIDWithGetOrCreate gets business unit ID using GetOrCreateBusinessUnitByDomainName
func (uc *UserController) getBusinessUnitIDWithGetOrCreate(ctx context.Context, userEmail string) string {
	domainName, isValid := utils.ExtractDomainFromEmail(userEmail)
	if !isValid {
		log.Warn().
			Str("user_email", userEmail).
			Msg("Invalid email format, cannot extract domain")
		return ""
	}

	// Sử dụng GetOrCreateBusinessUnitByDomainName thay vì logic phức tạp
	businessUnit, err := uc.services.BusinessUnit.GetOrCreateBusinessUnitByDomainName(ctx, domainName)
	if err != nil {
		log.Error().Err(err).
			Str("domain_name", domainName).
			Str("user_email", userEmail).
			Msg("Failed to get or create business unit")
		return ""
	}

	log.Info().
		Str("domain_name", domainName).
		Str("business_unit_id", businessUnit.ID).
		Str("business_unit_name", businessUnit.Name).
		Str("user_email", userEmail).
		Msg("Successfully got or created business unit")

	return businessUnit.ID
}

// buildUserInfoResponse creates the user info response object
func (uc *UserController) buildUserInfoResponse(user models.Userable, businessUnitName string) responseModel.UserInfoResponse {
	return responseModel.UserInfoResponse{
		ID:             utils.GetStringValue(user.GetId()),
		DisplayName:    utils.GetStringValue(user.GetDisplayName()),
		Surname:        utils.GetStringValue(user.GetSurname()),
		GivenName:      utils.GetStringValue(user.GetGivenName()),
		Email:          utils.GetStringValue(user.GetMail()),
		MobilePhone:    utils.GetStringValue(user.GetMobilePhone()),
		JobTitle:       utils.GetStringValue(user.GetJobTitle()),
		OfficeLocation: utils.GetStringValue(user.GetOfficeLocation()),
		Department:     utils.GetStringValue(user.GetDepartment()),
		BusinessUnit:   businessUnitName,
		Manager:        responseModel.GetManagerId(user),
		Status:         model.UserStatusActive,
	}
}
