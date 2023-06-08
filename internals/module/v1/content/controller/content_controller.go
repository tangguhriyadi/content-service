package controller

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/service"
	"github.com/tangguhriyadi/content-service/internals/security/token"
)

type ContentController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type ContentControllerImpl struct {
	validate       *validator.Validate
	contentService service.ContentService
}

func NewContentController(validate *validator.Validate, contentService service.ContentService) ContentController {
	return &ContentControllerImpl{
		validate:       validate,
		contentService: contentService,
	}
}

// ShowAccount godoc
// @Summary      get content list
// @Description  get content list
// @Tags         contents
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.ContentPaginate
// @Router       /contents [get]
// @Security 	 Bearer
func (cc ContentControllerImpl) GetAll(ctx *fiber.Ctx) error {
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

	result, err := cc.contentService.GetAll(c, page, limit)

	if result.Data == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to retrieve data",
		})
	}

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      create content
// @Description  create content
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param payload body dto.ContentPayload true "The input struct"
// @Success      200  {object}  dto.Content
// @Router       /contents/:id [post]
// @Security 	 Bearer
func (cc ContentControllerImpl) Create(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var payload dto.ContentPayload

	//body parsing
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// request body validation
	if err := cc.validate.Struct(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// claim token
	claims, err := token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	//run business logic
	if err := cc.contentService.Create(c, payload, int32(claims.UserId)); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}
	fmt.Println("masuk4")
	return helper.ApiResponse(ctx, true, "Create Success", "", &payload, fiber.StatusCreated)

}

// ShowAccount godoc
// @Summary      get content by ID
// @Description  get content by ID
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param id path string true "Content ID"
// @Success      200  {object}  dto.Content
// @Router       /contents/:id [get]
// @Security 	 Bearer
func (cc ContentControllerImpl) GetById(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var userId = ctx.Params("id")

	// param validator
	user_id, err := strconv.Atoi(userId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//run business logic
	result, err := cc.contentService.GetById(c, user_id)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Get", "", result, fiber.StatusOK)
}

// ShowAccount godoc
// @Summary      update content
// @Description  update content
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param id path string true "Content ID"
// @Success      200  {object}  dto.ContentPayload
// @Router       /contents/:id [patch]
// @Security 	 Bearer
func (cc ContentControllerImpl) Update(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var userId = ctx.Params("id")
	var payload dto.ContentPayload

	// param validator
	user_id, err := strconv.Atoi(userId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//body parsing
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// run business logic
	res := cc.contentService.Update(c, user_id, &payload)
	if res != nil {
		return helper.ApiResponse(ctx, false, "Internal Server Error", res.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Update", "", nil, fiber.StatusOK)
}

// ShowAccount godoc
// @Summary      delete content
// @Description  delete content
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param id path string true "Content ID"
// @Success      200  {object}  dto.Content
// @Router       /contents/:id [delete]
// @Security 	 Bearer
func (cc ContentControllerImpl) Delete(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var userId = ctx.Params("id")

	// param validator
	user_id, err := strconv.Atoi(userId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// run business logic
	res := cc.contentService.Delete(c, user_id)
	if res != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", res.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Delete", "", nil, fiber.StatusOK)
}
