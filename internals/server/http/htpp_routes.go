package http

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content"
	contentType "github.com/tangguhriyadi/content-service/internals/module/v1/content_type"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func HttpRouteInit(r *fiber.App, containerConf *container.Container) {

	contentApi := r.Group("/dapur")
	content.ContentRoutes(contentApi, containerConf)

	contentTypeApi := r.Group("/types")
	contentType.ContentTypeRoutes(contentTypeApi, containerConf)

	r.Get("/documentation/*", fiberSwagger.WrapHandler)
	r.Use(func(ctx *fiber.Ctx) error {
		return helper.ApiResponse(ctx, true, "NOT FOUND", "", nil, fiber.StatusOK)
	})

}
