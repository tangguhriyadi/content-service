package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/helper"
)

func HttpRouteInit(r *fiber.App) {
	// api := r.Group("/v1")

	// contentApi := api.Group("/contents")

	r.Use(func(ctx *fiber.Ctx) error {
		return helper.ApiResponse(ctx, true, "NOT FOUND", "", nil, fiber.StatusOK)
	})

}
