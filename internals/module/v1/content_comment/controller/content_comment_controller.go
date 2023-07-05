package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/service"
	"github.com/tangguhriyadi/content-service/internals/security/token"
)

type ContentCommentController interface {
	GetByContentId(ctx *fiber.Ctx) error
	PostComment(ctx *fiber.Ctx) error
}

type ContentCommentControllerImpl struct {
	validate              *validator.Validate
	contentCommentService service.ContentCommentService
}

func NewContentCommentController(contentCommentService service.ContentCommentService, validate *validator.Validate) ContentCommentController {
	return &ContentCommentControllerImpl{
		validate:              validate,
		contentCommentService: contentCommentService,
	}
}

// ShowAccount godoc
// @Summary      get content comment by id
// @Description  get content comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param	id path	string	false	"id"
// @Success      200  {object}  []dto.ContentComment
// @Router       /contents/:id/comment/ [get]
// @Security 	 Bearer
func (cc ContentCommentControllerImpl) GetByContentId(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var contentId = ctx.Params("id")

	// param validator
	content_id, err := strconv.Atoi(contentId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	//run business logic
	result, err := cc.contentCommentService.GetByContentId(c, int32(content_id))
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Success Get", "", result, fiber.StatusOK)
}

// ShowAccount godoc
// @Summary      create comment
// @Description  create comment
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param payload body dto.ContentPayload true "The input struct"
// @Success      200  {object}  dto.Content
// @Router       /contents/:id/comment [post]
// @Security 	 Bearer
func (cc ContentCommentControllerImpl) PostComment(ctx *fiber.Ctx) error {
	var c = ctx.Context()
	var contentId = ctx.Params("id")
	var payload dto.ContentPayload

	// param validator
	content_id, err := strconv.Atoi(contentId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// body parsing
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	// request body validation
	if err := cc.validate.Struct(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	token, err := token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	if err := cc.contentCommentService.PostComment(c, int32(content_id), &payload, int32(token.UserId)); err != nil {
		return helper.ApiResponse(ctx, true, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ApiResponse(ctx, true, "Comment Success", "", &payload, fiber.StatusCreated)
}
