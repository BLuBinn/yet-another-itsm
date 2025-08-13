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

type FormSectionController struct {
	services *service.Services
}

func NewFormSectionController(services *service.Services) *FormSectionController {
	return &FormSectionController{
		services: services,
	}
}

// GetFormSections godoc
// @Summary Get all form sections
// @Description Get all form sections
// @Tags form-sections
// @Accept json
// @Produce json
// @Success 200 {object} responseModel.FormSectionsListResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-sections [get]
func (fs *FormSectionController) GetFormSections(c *gin.Context) {
	log.Info().
		Str("controller", "FormSectionController").
		Str("endpoint", "GetFormSections").
		Str("method", c.Request.Method).
		Msg("Get all form sections endpoint called")

	ctx := c.Request.Context()
	templateID := c.Query("templateId")

	sections, err := fs.services.FormSection.GetFormSections(ctx, templateID)
	if err != nil {
		log.Error().Err(err).Str("templateId", templateID).Msg(constants.ErrFailedToGetFormSections)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormSections)
		return
	}

	response := &responseModel.FormSectionsListResponse{
		Items: make([]responseModel.FormSectionResponse, len(sections)),
	}

	for i, section := range sections {
		response.Items[i] = *section.ToResponse()
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessGetFormSections, response)
}

// GetFormSectionByID godoc
// @Summary Get form section by ID
// @Description Get form section by ID
// @Tags form-sections
// @Accept json
// @Produce json
// @Param sectionId path string true "Section ID"
// @Success 200 {object} responseModel.FormSectionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-sections/{sectionId} [get]
func (fs *FormSectionController) GetFormSectionByID(c *gin.Context) {
	log.Info().
		Str("controller", "FormSectionController").
		Str("endpoint", "GetFormSectionByID").
		Str("method", c.Request.Method).
		Msg("Get form section by ID endpoint called")

	sectionID := c.Param("sectionId")
	ctx := c.Request.Context()

	section, err := fs.services.FormSection.GetFormSectionByID(ctx, sectionID)
	if err != nil {
		log.Error().Err(err).Str("sectionId", sectionID).Msg(constants.ErrFailedToGetFormSection)
		utils.SendInternalServerError(c, constants.ErrFailedToGetFormSection)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessGetFormSections, section.ToResponse())
}

// CreateFormSection godoc
// @Summary Create form section
// @Description Create a new form section
// @Tags form-sections
// @Accept json
// @Produce json
// @Param request body responseModel.CreateFormSectionRequest true "Create form section request"
// @Success 201 {object} responseModel.FormSectionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-sections [post]
func (fs *FormSectionController) CreateFormSection(c *gin.Context) {
	log.Info().
		Str("controller", "FormSectionController").
		Str("endpoint", "CreateFormSection").
		Str("method", c.Request.Method).
		Msg("Create form section endpoint called")

	var req responseModel.CreateFormSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBody)
		return
	}

	ctx := c.Request.Context()

	section, err := fs.services.FormSection.CreateFormSection(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg(constants.ErrFailedToCreateFormSection)
		utils.SendInternalServerError(c, constants.ErrFailedToCreateFormSection)
		return
	}

	utils.SendSuccess(c, http.StatusCreated, constants.SuccessCreateFormSection, section.ToResponse())
}

// UpdateFormSection godoc
// @Summary Update form section
// @Description Update an existing form section
// @Tags form-sections
// @Accept json
// @Produce json
// @Param sectionId path string true "Section ID"
// @Param request body responseModel.UpdateFormSectionRequest true "Update form section request"
// @Success 200 {object} responseModel.FormSectionResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-sections/{sectionId} [put]
func (fs *FormSectionController) UpdateFormSection(c *gin.Context) {
	log.Info().
		Str("controller", "FormSectionController").
		Str("endpoint", "UpdateFormSection").
		Str("method", c.Request.Method).
		Msg("Update form section endpoint called")

	sectionID := c.Param("sectionId")

	var req responseModel.UpdateFormSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		utils.SendBadRequest(c, constants.ErrInvalidRequestBody)
		return
	}

	ctx := c.Request.Context()

	section, err := fs.services.FormSection.UpdateFormSection(ctx, sectionID, &req)
	if err != nil {
		log.Error().Err(err).Str("sectionId", sectionID).Msg(constants.ErrFailedToUpdateFormSection)
		utils.SendInternalServerError(c, constants.ErrFailedToUpdateFormSection)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessUpdateFormSection, section.ToResponse())
}

// DeleteFormSection godoc
// @Summary Delete form section
// @Description Delete a form section
// @Tags form-sections
// @Accept json
// @Produce json
// @Param sectionId path string true "Section ID"
// @Success 200 {object} responseModel.SuccessResponse
// @Failure 400 {object} responseModel.ErrorResponse
// @Failure 404 {object} responseModel.ErrorResponse
// @Failure 500 {object} responseModel.ErrorResponse
// @Router /v1/form-sections/{sectionId} [delete]
func (fs *FormSectionController) DeleteFormSection(c *gin.Context) {
	log.Info().
		Str("controller", "FormSectionController").
		Str("endpoint", "DeleteFormSection").
		Str("method", c.Request.Method).
		Msg("Delete form section endpoint called")

	sectionID := c.Param("sectionId")
	ctx := c.Request.Context()

	err := fs.services.FormSection.DeleteFormSection(ctx, sectionID)
	if err != nil {
		log.Error().Err(err).Str("sectionId", sectionID).Msg(constants.ErrFailedToDeleteFormSection)
		utils.SendInternalServerError(c, constants.ErrFailedToDeleteFormSection)
		return
	}

	utils.SendSuccess(c, http.StatusOK, constants.SuccessDeleteFormSection, nil)
}
