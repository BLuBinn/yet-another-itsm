package controller

import (
	"net/http"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/service"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
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
	existingUser, err := uc.services.User.GetUserByEmail(ctx, userEmail)

	if err != nil || existingUser == nil {
		log.Info().
			Str("email", userEmail).
			Msg("User not found in database, creating new user")

		createUserReq := &responseModel.CreateUserRequest{
			AzureAdObjectID: utils.GetStringValue(user.GetId()),
			// HomeTenantID:    utils.GetStringValue(user.GetTenantId()),
			Mail:           userEmail,
			DisplayName:    utils.GetStringValue(user.GetDisplayName()),
			GivenName:      utils.GetStringValue(user.GetGivenName()),
			SurName:        utils.GetStringValue(user.GetSurname()),
			JobTitle:       utils.GetStringValue(user.GetJobTitle()),
			OfficeLocation: utils.GetStringValue(user.GetOfficeLocation()),
			DepartmentID:   utils.GetStringValue(user.GetDepartment()),
			ManagerID:      responseModel.GetManagerId(user),
			IsActive:       true,
		}

		_, createErr := uc.services.User.CreateUser(ctx, createUserReq)
		if createErr != nil {
			log.Error().Err(createErr).Msg("Failed to create user in database")
		} else {
			log.Info().
				Str("email", userEmail).
				Msg("Successfully created user in database")
		}
	}

	UserInfoResponse := responseModel.UserInfoResponse{
		ID:             utils.GetStringValue(user.GetId()),
		DisplayName:    utils.GetStringValue(user.GetDisplayName()),
		Surname:        utils.GetStringValue(user.GetSurname()),
		GivenName:      utils.GetStringValue(user.GetGivenName()),
		Email:          utils.GetStringValue(user.GetMail()),
		MobilePhone:    utils.GetStringValue(user.GetMobilePhone()),
		JobTitle:       utils.GetStringValue(user.GetJobTitle()),
		OfficeLocation: utils.GetStringValue(user.GetOfficeLocation()),
		Department:     utils.GetStringValue(user.GetDepartment()),
		Manager:        responseModel.GetManagerId(user),
	}

	log.Info().
		Str("user_id", UserInfoResponse.ID).
		Msg(constants.SuccessMsgGetCurrentUser)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsgGetCurrentUser, UserInfoResponse)
}
