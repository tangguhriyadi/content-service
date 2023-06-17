package contentlike

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/controller"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/service"
	"github.com/tangguhriyadi/content-service/internals/security/middleware"
)

func ContentLikeRoutes(r fiber.Router, containerConf *container.Container) {
	repository := repository.NewContentLikeRepo(containerConf.Postgre)
	service := service.NewContentLikeService(repository)
	controller := controller.NewContentLikeController(containerConf.Validator, service)

	r.Post("/:id/action", middleware.JWTProtect(), controller.Like)
}
