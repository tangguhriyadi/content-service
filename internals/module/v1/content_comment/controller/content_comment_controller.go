package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/service"
)

type ContentCommentController interface {
	GetByContentId(ctx *fiber.Ctx) error
}

type ContentCommentControllerImpl struct {
	contentCommentService service.ContentCommentService
}

func NewContentCommentController(contentCommentService service.ContentCommentService) ContentCommentController {
	return &ContentCommentControllerImpl{
		contentCommentService: contentCommentService,
	}
}

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
