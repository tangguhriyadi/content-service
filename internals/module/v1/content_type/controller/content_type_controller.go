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
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param	page query	string	false	"page"
// @Param	limit query	string	false	"limit page"
// @Success      200  {object}  dto.ContentTypePaginate
// @Router       /contents/types [get]
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
