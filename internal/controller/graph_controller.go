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

type GraphController struct {
	services *service.Services
}

func NewGraphController(services *service.Services) *GraphController {
	return &GraphController{
		services: services,
	}
}

// GetCurrentUser godoc
// @Summary Get user information
// @Description Get user information
// @Tags graph
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserInfoResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /v1/graph/users/me [get]
func (gc *GraphController) GetCurrentUserFromGraph(c *gin.Context) {
	log.Info().
		Str("controller", "GraphController").
		Str("endpoint", "GetCurrentUserFromGraph").
		Str("method", c.Request.Method).
		Msg("Getting current user from graph")

	accessToken, err := utils.GetAccessToken(c.Request.Context())
	if err != nil {
		utils.SendUnauthorized(c, constants.ErrAccessTokenNotFoundMsg)
		return
	}

	user, err := gc.services.Graph.GetCurrentUserFromGraph(accessToken)
	if err != nil {
		utils.SendUnauthorized(c, constants.ErrFailedToGetUserInformationMsg)
		return
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
