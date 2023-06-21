package http

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content"
	contentComment "github.com/tangguhriyadi/content-service/internals/module/v1/content_comment"
	contentLike "github.com/tangguhriyadi/content-service/internals/module/v1/content_like"
	contentType "github.com/tangguhriyadi/content-service/internals/module/v1/content_type"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func HttpRouteInit(r *fiber.App, containerConf *container.Container) {

	contentApi := r.Group("/")
	content.ContentRoutes(contentApi, containerConf)
	contentType.ContentTypeRoutes(contentApi, containerConf)
	contentLike.ContentLikeRoutes(contentApi, containerConf)
	contentComment.ContentCommentRoutes(contentApi, containerConf)

	r.Get("/documentation/*", fiberSwagger.WrapHandler)
	r.Use(func(ctx *fiber.Ctx) error {
		return helper.ApiResponse(ctx, false, "NOT FOUND", "", nil, fiber.StatusOK)
	})

}
