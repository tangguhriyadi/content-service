package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content"
)

func HttpRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/v1")

	contentApi := api.Group("/contents")
	content.ContentRoutes(contentApi, containerConf)

	r.Use(func(ctx *fiber.Ctx) error {
		return helper.ApiResponse(ctx, true, "NOT FOUND", "", nil, fiber.StatusOK)
	})

}
