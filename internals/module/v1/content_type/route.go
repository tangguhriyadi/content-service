package contenttype

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/controller"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/service"
	"github.com/tangguhriyadi/content-service/internals/security/middleware"
)

func ContentTypeRoutes(r fiber.Router, containerConf *container.Container) {
	repository := repository.NewContentTypeRepository(containerConf.Postgre)
	service := service.NewContentTypeService(repository)
	controller := controller.NewContentTypeController(containerConf.Validator, service)

	r.Get("/", middleware.JWTProtect(), controller.GetAll)
}
