package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/service"
	"github.com/tangguhriyadi/content-service/internals/security/token"
)

type ContentLikeController interface {
	Like(ctx *fiber.Ctx) error
}

type ContentLikeControllerImpl struct {
	validate           *validator.Validate
	contentLikeService service.ContentLikeService
}

func NewContentLikeController(validate *validator.Validate,
	contentLikeService service.ContentLikeService) ContentLikeController {
	return &ContentLikeControllerImpl{
		validate:           validate,
		contentLikeService: contentLikeService,
	}
}

// ShowAccount godoc
// @Summary      post content like
// @Description  post content like
// @Tags         contents
// @Accept       json
// @Produce      json
// @Param	id path	string	true	"content id"
// @Param	payload body dto.ContentLikePayload	true "type"
// @Success      200  {object}  dto.ContentLikePayload
// @Router       /contents/:id/like [post]
// @Security 	 Bearer
func (cl ContentLikeControllerImpl) Like(ctx *fiber.Ctx) error {
	c := ctx.Context()
	var payload dto.ContentLikePayload
	var contentId = ctx.Params("id")

	content_id, err := strconv.Atoi(contentId)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	if err := cl.validate.Struct(&payload); err != nil {
		return helper.ApiResponse(ctx, false, "Bad Request", err.Error(), nil, fiber.StatusBadRequest)
	}

	token, err := token.ExtractTokenMetada(ctx)
	if err != nil {
		return helper.ApiResponse(ctx, false, "Forbidden", err.Error(), nil, fiber.StatusForbidden)
	}

	if err := cl.contentLikeService.Like(c, payload.Type, int32(content_id), int32(token.UserId)); err != nil {
		return helper.ApiResponse(ctx, false, "Internal Server Error", err.Error(), nil, fiber.StatusInternalServerError)
	}

	return helper.ApiResponse(ctx, true, "Content Status Updated !", "", nil, fiber.StatusOK)

}
