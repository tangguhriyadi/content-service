package contentlike

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	contentRepo "github.com/tangguhriyadi/content-service/internals/module/v1/content/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/controller"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/service"
	"github.com/tangguhriyadi/content-service/internals/security/middleware"
)

func ContentLikeRoutes(r fiber.Router, containerConf *container.Container) {
	contentRepo := contentRepo.NewContentRepository(containerConf.Postgre)
	repository := repository.NewContentLikeRepo(containerConf.Postgre)
	service := service.NewContentLikeService(repository, contentRepo)
	controller := controller.NewContentLikeController(containerConf.Validator, service)

	r.Post("/:id/action", middleware.JWTProtect(), controller.Like)
}
