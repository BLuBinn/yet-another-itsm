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

type ScopeController struct {
	services *service.Services
}

func NewScopeController(services *service.Services) *ScopeController {
	return &ScopeController{
		services: services,
	}
}

// GetAllScopes godoc
// @Summary Get all scopes
// @Description Get all scopes in the system
// @Tags scopes
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.ScopesListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/scopes [get]
func (sc *ScopeController) GetAllScopes(c *gin.Context) {
	log.Info().
		Str("controller", "ScopeController").
		Str("endpoint", "GetAllScopes").
		Str("method", c.Request.Method).
		Msg("Get all scopes endpoint called")

	ctx := c.Request.Context()

	scopes, err := sc.services.Scope.GetAllScopes(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get scopes")
		utils.SendInternalServerError(c, constants.ErrFailedToRetrieveScopesMsg)
		return
	}

	log.Info().
		Int("count", len(scopes)).
		Msg("Successfully retrieved scopes")

	var scopeResponses []responseModel.ScopeResponse
	for _, scope := range scopes {
		scopeResponses = append(scopeResponses, *scope.ToResponse())
	}

	responseData := responseModel.NewScopesListResponse(
		scopeResponses,
		1,
		len(scopeResponses),
		int64(len(scopeResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved scopes", responseData)
}

// GetScopeByID godoc
// @Summary Get scope by ID
// @Description Get a specific scope by ID
// @Tags scopes
// @Accept json
// @Produce json
// @Param scopeId path string true "Scope ID"
// @Success 200 {object} responseModel.ScopeResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/scopes/{scopeId} [get]
func (sc *ScopeController) GetScopeByID(c *gin.Context) {
	log.Info().
		Str("controller", "ScopeController").
		Str("endpoint", "GetScopeByID").
		Str("method", c.Request.Method).
		Msg("Get scope by ID endpoint called")

	id := c.Param("scopeId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrScopeIDRequiredMsg)
		return
	}

	ctx := c.Request.Context()

	scope, err := sc.services.Scope.GetScopeByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("scope_id", id).Msg("Failed to get scope by ID")
		utils.SendNotFound(c, constants.ErrScopeNotFoundMsg)
		return
	}

	log.Info().
		Str("scope_id", id).
		Str("scope_name", scope.Name).
		Msg("Successfully retrieved scope by ID")

	utils.SendSuccess(c, http.StatusOK, "Successfully retrieved scope by ID", scope.ToResponse())
}

// CreateScope godoc
// @Summary Create a new scope
// @Description Create a new scope
// @Tags scopes
// @Accept json
// @Produce json
// @Param scope body responseModel.CreateScopeRequest true "Scope data"
// @Success 201 {object} responseModel.ScopeResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/scopes [post]
func (sc *ScopeController) CreateScope(c *gin.Context) {
	var req responseModel.CreateScopeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request body")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	log.Info().
		Str("controller", "ScopeController").
		Str("method", "CreateScope").
		Str("scope_id", req.ID).
		Str("scope_name", req.Name).
		Msg("Creating new scope")

	ctx := c.Request.Context()

	scope, err := sc.services.Scope.CreateScope(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create scope")
		utils.SendInternalServerError(c, constants.ErrFailedToCreateScopeMsg)
		return
	}

	log.Info().
		Str("scope_id", scope.ID).
		Str("scope_name", scope.Name).
		Msg("Successfully created scope")

	utils.SendSuccess(c, http.StatusCreated, "Successfully created scope", scope.ToResponse())
}
