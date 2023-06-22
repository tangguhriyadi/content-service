package contentcomment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/controller"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/service"
	"github.com/tangguhriyadi/content-service/internals/security/middleware"
)

func ContentCommentRoutes(r fiber.Router, containerConf *container.Container) {
	repository := repository.NewContentCommentRepo(containerConf.Postgre)
	service := service.NewContentCommentService(repository)
	controller := controller.NewContentCommentController(service, containerConf.Validator)

	r.Get("/:id/comment", middleware.JWTProtect(), controller.GetByContentId)
	r.Post("/:id/comment", middleware.JWTProtect(), controller.PostComment)
}
