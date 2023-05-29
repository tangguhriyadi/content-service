package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/tangguhriyadi/content-service/internals/config"
	"github.com/tangguhriyadi/content-service/internals/helper"
)

func JWTProtect() func(*fiber.Ctx) error {
	env := config.New()
	config := jwtMiddleware.Config{
		SigningKey:   []byte(env.Get("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(ctx *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return helper.ApiResponse(ctx, false, "Unauthorized", "Unauthorized", nil, fiber.StatusUnauthorized)
}
