package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/service"
)

type ContentTypeController interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
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
