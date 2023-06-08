package content

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/controller"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/service"
	"github.com/tangguhriyadi/content-service/internals/security/middleware"
)

func ContentRoutes(r fiber.Router, containerConf *container.Container) {
	repository := repository.NewContentRepository(containerConf.Postgre)
	service := service.NewContentService(repository)
	controller := controller.NewContentController(containerConf.Validator, service)

	r.Get("/", middleware.JWTProtect(), controller.GetAll)
	r.Post("/", middleware.JWTProtect(), controller.Create)
	r.Get("/:id", middleware.JWTProtect(), controller.GetById)
	r.Patch("/:id", middleware.JWTProtect(), controller.Update)
	r.Delete("/:id", middleware.JWTProtect(), controller.Delete)
}
