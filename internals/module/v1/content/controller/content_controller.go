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
