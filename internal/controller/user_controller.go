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
}

func NewUserController(services *service.Services) *UserController {
	return &UserController{
		userService: services.User,
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
// @Router /v1/users/email/{email} [get]
func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserRequest true "User data"
// @Success 201 {object} dtos.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req responseModel.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg(constants.ErrInvalidRequestBodyMsg)
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "UserController").
		Str("endpoint", "CreateUser").
		Str("method", c.Request.Method).
		Str("email", req.Mail).
		Msg("Creating new user")

	ctx := c.Request.Context()

	user, err := uc.userService.CreateUser(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		utils.SendInternalServerError(c, constants.ErrFailedToCreateUserMsg)
		return
	}

	utils.SendSuccess(c, http.StatusCreated, constants.SuccessMsgCreateUser, user.ToResponse())
}
