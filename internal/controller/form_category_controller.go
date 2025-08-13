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

type FormCategoryController struct {
	services *service.Services
}

func NewFormCategoryController(services *service.Services) *FormCategoryController {
	return &FormCategoryController{
		services: services,
	}
}

// GetFormCategories godoc
// @Summary Get all form categories
// @Description Get all form categories
// @Tags form-categories
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.FormCategoriesListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-categories [get]
func (fc *FormCategoryController) GetFormCategories(c *gin.Context) {
	log.Info().
		Str("controller", "FormCategoryController").
		Str("endpoint", "GetFormCategories").
		Str("method", c.Request.Method).
		Msg("Get all form categories endpoint called")

	ctx := c.Request.Context()

	categories, err := fc.services.FormCategory.GetFormCategories(ctx)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToGetFormCategories)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormCategories)
		return
	}

	var categoryResponses []responseModel.FormCategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, *category.ToResponse())
	}

	response := responseModel.NewFormCategoriesListResponse(
		categoryResponses,
		1,
		len(categoryResponses),
		int64(len(categoryResponses)),
	)

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsg, response)
}

// GetFormCategoryByID godoc
// @Summary Get form category by ID
// @Description Get form category by ID
// @Tags form-categories
// @Accept json
// @Produce json
// @Param categoryId path string true "Form Category ID"
// @Success 200 {object} responseModel.FormCategoryResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-categories/{categoryId} [get]
func (fc *FormCategoryController) GetFormCategoryByID(c *gin.Context) {
	log.Info().
		Str("controller", "FormCategoryController").
		Str("endpoint", "GetFormCategoryByID").
		Str("method", c.Request.Method).
		Msg("Get form category by ID endpoint called")

	id := c.Param("categoryId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrFormCategoryIDRequired)
		return
	}

	ctx := c.Request.Context()

	category, err := fc.services.FormCategory.GetFormCategoryByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToGetFormCategory)
		utils.SendNotFound(c, constants.ErrFormCategoryNotFound)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsg, category.ToResponse())
}

// CreateFormCategory godoc
// @Summary Create a new form category
// @Description Create a new form category
// @Tags form-categories
// @Accept json
// @Produce json
// @Param request body responseModel.CreateFormCategoryRequest true "Create form category request"
// @Success 201 {object} responseModel.FormCategoryResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-categories [post]
func (fc *FormCategoryController) CreateFormCategory(c *gin.Context) {
	log.Info().
		Str("controller", "FormCategoryController").
		Str("endpoint", "CreateFormCategory").
		Str("method", c.Request.Method).
		Msg("Create form category endpoint called")

	var request responseModel.CreateFormCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	ctx := c.Request.Context()

	category, err := fc.services.FormCategory.CreateFormCategory(ctx, request.Name, request.Description)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToCreateFormCategory)
		utils.SendInternalServerError(c, constants.ErrFailedToCreateFormCategory)
		return
	}

	utils.SendSuccess(c, http.StatusCreated, constants.SuccessMsg, category.ToResponse())
}

// UpdateFormCategory godoc
// @Summary Update a form category
// @Description Update a form category
// @Tags form-categories
// @Accept json
// @Produce json
// @Param categoryId path string true "Form Category ID"
// @Param request body responseModel.UpdateFormCategoryRequest true "Update form category request"
// @Success 200 {object} responseModel.FormCategoryResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-categories/{categoryId} [put]
func (fc *FormCategoryController) UpdateFormCategory(c *gin.Context) {
	log.Info().
		Str("controller", "FormCategoryController").
		Str("endpoint", "UpdateFormCategory").
		Str("method", c.Request.Method).
		Msg("Update form category endpoint called")

	id := c.Param("categoryId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrFormCategoryIDRequired)
		return
	}

	var request responseModel.UpdateFormCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendBadRequest(c, constants.ErrInvalidRequestBodyMsg)
		return
	}

	ctx := c.Request.Context()

	category, err := fc.services.FormCategory.UpdateFormCategory(ctx, id, request.Name, request.Description)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToUpdateFormCategory)
		utils.SendNotFound(c, constants.ErrFormCategoryNotFound)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsg, category.ToResponse())
}

// DeleteFormCategory godoc
// @Summary Delete a form category
// @Description Delete a form category
// @Tags form-categories
// @Accept json
// @Produce json
// @Param categoryId path string true "Form Category ID"
// @Success 200 {object} responseModel.SuccessResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-categories/{categoryId} [delete]
func (fc *FormCategoryController) DeleteFormCategory(c *gin.Context) {
	log.Info().
		Str("controller", "FormCategoryController").
		Str("endpoint", "DeleteFormCategory").
		Str("method", c.Request.Method).
		Msg("Delete form category endpoint called")

	id := c.Param("categoryId")
	if id == "" {
		utils.SendBadRequest(c, constants.ErrFormCategoryIDRequired)
		return
	}

	ctx := c.Request.Context()

	err := fc.services.FormCategory.DeleteFormCategory(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg(constants.ErrFailedToDeleteFormCategory)
		utils.SendNotFound(c, constants.ErrFormCategoryNotFound)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessMsg, nil)
}
