package http

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func HttpRouteInit(r *fiber.App, containerConf *container.Container) {

	r.Get("/documentation/contents/*", fiberSwagger.WrapHandler)

	contentApi := r.Group("/")
	content.ContentRoutes(contentApi, containerConf)

	r.Use(func(ctx *fiber.Ctx) error {
		return helper.ApiResponse(ctx, true, "NOT FOUND", "", nil, fiber.StatusOK)
	})

}
