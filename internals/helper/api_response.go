package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
	Data       interface{} `json:"data"`
}

func ApiResponse(ctx *fiber.Ctx, status bool, message string, err string, data interface{}, code int) error {
	var splittedError []string

	if len(err) > 0 {
		splittedError = strings.Split(err, "\n")
	}

	return ctx.Status(code).JSON(&Response{
		Success:    status,
		StatusCode: code,
		Message:    message,
		Errors:     splittedError,
		Data:       data,
	})
}
