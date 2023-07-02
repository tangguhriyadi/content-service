package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/service"
	"github.com/tangguhriyadi/content-service/internals/security/token"
)

type ContentTypeController interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type ContentTypeControllerImpl struct {
	validate           *validator.Validate
	contentTypeService service.ContentTypeService
}

func NewContentTypeController(validate *validator.Validate, contentTypeService service.ContentTypeService) ContentTypeController {
	return &ContentTypeControllerImpl{
		validate:           validate,
		contentTypeService: contentTypeService,
	}
}

// ShowAccount godoc
// @Summary      get content type list
// @Description  get content type list
// @Tags         types
// @Accept       json
// @Produce      json
// @Param	page query	string	false	"page"
// @Param	limit query	string	false	"limit page"
// @Success      200  {object}  dto.ContentTypePaginate
// @Router       /contents/:id/types [get]
// @Security 	 Bearer
func (ct ContentTypeControllerImpl) GetAll(ctx *fiber.Ctx) error {
	c := ctx.Context()

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "query page must be int",
		})
	}
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "query limit must be int",
		})
	}
	result, err := ct.contentTypeService.GetAll(c, page, limit)

	if result.Data == nil {
		return helper.ApiResponse(ctx, false, "Bad Request", "failed to retrieve data", nil, fiber.StatusBadRequest)
	}

	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Get List", "", &result, fiber.StatusOK)
}

// ShowAccount godoc
// @Summary      get content type by id
// @Description  get content type by id
// @Tags         types
// @Accept       json
// @Produce      json
// @Param	id path	string	false	"id"
// @Param	type_id path	string	false	"type_id"
// @Success      200  {object}  dto.ContentTypePaginate
// @Router       /contents/:id/types/:type_id [get]
// @Security 	 Bearer
func (ct ContentTypeControllerImpl) GetById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	var id = ctx.Params("type_id")

	//param validator
	typeId, err := strconv.Atoi(id)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//run business logic
	result, err := ct.contentTypeService.GetById(c, typeId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	if result.ID == 0 {
		return helper.ApiResponse(ctx, false, "No Data Found", "", nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Get By Id", "", result, fiber.StatusOK)

}

// ShowAccount godoc
// @Summary      create content type by id
// @Description  create content type by id
// @Tags         types
// @Accept       json
// @Produce      json
// @Param	id path	string	false	"id"
// @Param payload body dto.ContentTypePayload true "The input struct"
// @Success      200  {object}  dto.ContentTypePayload
// @Router       /contents/:id/types [post]
// @Security 	 Bearer
func (ct ContentTypeControllerImpl) Create(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var payload dto.ContentTypePayload

	//body parsing
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//request body validation
	if err := ct.validate.Struct(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//claim token
	token, err := token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	if err := ct.contentTypeService.Create(c, &payload, int32(token.UserId)); err != nil {
		return helper.ApiResponse(ctx, true, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Create Success", "", &payload, fiber.StatusCreated)
}

// ShowAccount godoc
// @Summary      update content type by id
// @Description  update content type by id
// @Tags         types
// @Accept       json
// @Produce      json
// @Param	id path	string	false	"id"
// @param type_id path string false "type_id"
// @Param payload body dto.ContentTypePayload true "The input struct"
// @Success      200  {object}  dto.ContentTypePayload
// @Router       /contents/:id/types [patch]
// @Security 	 Bearer
func (ct ContentTypeControllerImpl) Update(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var payload dto.ContentTypePayload
	var id = ctx.Params("type_id")

	typeId, err := strconv.Atoi(id)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}
	//body parsing
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}
	//request body validation
	if err := ct.validate.Struct(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}
	//claim token
	_, err = token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	if err := ct.contentTypeService.Update(c, typeId, &payload); err != nil {
		return helper.ApiResponse(ctx, true, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Update Success", "", &payload, fiber.StatusCreated)
}

func (ct ContentTypeControllerImpl) Delete(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var id = ctx.Params("type_id")
	typeId, err := strconv.Atoi(id)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//claim token
	token, err := token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	if err := ct.contentTypeService.Delete(c, typeId, int32(token.UserId)); err != nil {
		return helper.ApiResponse(ctx, true, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Delete Success", "", "", fiber.StatusOK)
}
