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

type FormTemplateController struct {
	services *service.Services
}

func NewFormTemplateController(services *service.Services) *FormTemplateController {
	return &FormTemplateController{
		services: services,
	}
}

// GetFormTemplates godoc
// @Summary Get all form templates
// @Description Get all form templates
// @Tags form-templates
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.FormTemplatesListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates [get]
func (ft *FormTemplateController) GetFormTemplates(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "GetFormTemplates").
		Str("method", c.Request.Method).
		Msg("Get all form templates endpoint called")

	ctx := c.Request.Context()

	templates, err := ft.services.FormTemplate.GetFormTemplates(ctx)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToGetFormTemplates)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormTemplates)
		return
	}

	response := &responseModel.FormTemplatesListResponse{
		Items: make([]responseModel.FormTemplateResponse, len(templates)),
	}

	for i, template := range templates {
		response.Items[i] = *template.ToResponse()
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessGetFormTemplates, response)
}

// GetFormTemplateByID godoc
// @Summary Get form template by ID
// @Description Get form template by ID
// @Tags form-templates
// @Accept json
// @Produce json
// @Param templateId path string true "Template ID"
// @Success 200 {object} responseModel.FormTemplateResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates/{templateId} [get]
func (ft *FormTemplateController) GetFormTemplateByID(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "GetFormTemplateByID").
		Str("method", c.Request.Method).
		Msg("Get form template by ID endpoint called")

	templateID := c.Param("templateId")
	ctx := c.Request.Context()

	template, err := ft.services.FormTemplate.GetFormTemplateByID(ctx, templateID)
	if err != nil {
		log.Error().Err(err).Str("templateId", templateID).Msg(constants.ErrFailedToGetFormTemplate)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormTemplate)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessGetFormTemplates, template.ToResponse())
}

// GetFormTemplatesByCategory godoc
// @Summary Get form templates by category
// @Description Get form templates by category ID
// @Tags form-templates
// @Accept json
// @Produce json
// @Param categoryId path string true "Category ID"
// @Success 200 {object} responseModel.FormTemplatesListResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates/category/{categoryId} [get]
func (ft *FormTemplateController) GetFormTemplatesByCategory(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "GetFormTemplatesByCategory").
		Str("method", c.Request.Method).
		Msg("Get form templates by category endpoint called")

	categoryID := c.Param("categoryId")
	ctx := c.Request.Context()

	templates, err := ft.services.FormTemplate.GetFormTemplatesByCategory(ctx, categoryID)
	if err != nil {
		log.Error().Err(err).Str("categoryId", categoryID).Msg(constants.ErrFailedToGetFormTemplates)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormTemplates)
		return
	}

	response := &responseModel.FormTemplatesListResponse{
		Items: make([]responseModel.FormTemplateResponse, len(templates)),
	}

	for i, template := range templates {
		response.Items[i] = *template.ToResponse()
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessGetFormTemplates, response)
}

// CreateFormTemplate godoc
// @Summary Create form template
// @Description Create a new form template
// @Tags form-templates
// @Accept json
// @Produce json
// @Param request body responseModel.CreateFormTemplateRequest true "Create form template request"
// @Success 201 {object} responseModel.FormTemplateResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates [post]
func (ft *FormTemplateController) CreateFormTemplate(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "CreateFormTemplate").
		Str("method", c.Request.Method).
		Msg("Create form template endpoint called")

	var req responseModel.CreateFormTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBody)
		return
	}

	ctx := c.Request.Context()

	template, err := ft.services.FormTemplate.CreateFormTemplate(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToCreateFormTemplate)
		utils.SendInternalServerError(c, constants.ErrFailedToCreateFormTemplate)
		return
	}

	utils.SendSuccess(c, http.StatusCreated, constants.SuccessCreateFormTemplate, template.ToResponse())
}

// UpdateFormTemplate godoc
// @Summary Update form template
// @Description Update an existing form template
// @Tags form-templates
// @Accept json
// @Produce json
// @Param templateId path string true "Template ID"
// @Param request body responseModel.UpdateFormTemplateRequest true "Update form template request"
// @Success 200 {object} responseModel.FormTemplateResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates/{templateId} [put]
func (ft *FormTemplateController) UpdateFormTemplate(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "UpdateFormTemplate").
		Str("method", c.Request.Method).
		Msg("Update form template endpoint called")

	templateID := c.Param("templateId")

	var req responseModel.UpdateFormTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBody)
		return
	}

	ctx := c.Request.Context()

	template, err := ft.services.FormTemplate.UpdateFormTemplate(ctx, templateID, &req)
	if err != nil {
		log.Error().Err(err).Str("templateId", templateID).Msg(constants.ErrFailedToUpdateFormTemplate)
		utils.SendInternalServerError(c, constants.ErrFailedToUpdateFormTemplate)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessUpdateFormTemplate, template.ToResponse())
}

// PublishFormTemplate godoc
// @Summary Publish form template
// @Description Publish a form template with version
// @Tags form-templates
// @Accept json
// @Produce json
// @Param templateId path string true "Template ID"
// @Param request body responseModel.PublishFormTemplateRequest true "Publish form template request"
// @Success 200 {object} responseModel.FormTemplateResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates/{templateId}/publish [post]
// func (ft *FormTemplateController) PublishFormTemplate(c *gin.Context) {
// 	log.Info().
// 		Str("controller", "FormTemplateController").
// 		Str("endpoint", "PublishFormTemplate").
// 		Str("method", c.Request.Method).
// 		Msg("Publish form template endpoint called")

// 	templateID := c.Param("templateId")

// 	var req responseModel.PublishFormTemplateRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		log.Error().Err(err).Msg("Failed to bind request")
// 		utils.SendBadRequestError(c, constants.ErrInvalidRequestBody)
// 		return
// 	}

// 	ctx := c.Request.Context()

// 	template, err := ft.services.FormTemplate.PublishFormTemplate(ctx, templateID, &req)
// 	if err != nil {
// 		log.Error().Err(err).Str("templateId", templateID).Msg(constants.ErrFailedToPublishFormTemplate)
// 		utils.SendInternalServerError(c, constants.ErrFailedToPublishFormTemplate)
// 		return
// 	}

// 	utils.SendSuccessResponse(c, http.StatusOK, constants.SuccessPublishFormTemplate, template.ToResponse())
// }

// DeleteFormTemplate godoc
// @Summary Delete form template
// @Description Delete a form template
// @Tags form-templates
// @Accept json
// @Produce json
// @Param templateId path string true "Template ID"
// @Success 200 {object} responseModel.SuccessResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-templates/{templateId} [delete]
func (ft *FormTemplateController) DeleteFormTemplate(c *gin.Context) {
	log.Info().
		Str("controller", "FormTemplateController").
		Str("endpoint", "DeleteFormTemplate").
		Str("method", c.Request.Method).
		Msg("Delete form template endpoint called")

	templateID := c.Param("templateId")
	ctx := c.Request.Context()

	err := ft.services.FormTemplate.DeleteFormTemplate(ctx, templateID)
	if err != nil {
		log.Error().Err(err).Str("templateId", templateID).Msg(constants.ErrFailedToDeleteFormTemplate)
		utils.SendInternalServerError(c, constants.ErrFailedToDeleteFormTemplate)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessDeleteFormTemplate, nil)
}
