package service

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/repository"
)

type ContentLikeService interface {
	Like(c context.Context, types string, content_id int32, user_id int32) error
}

type ContentLikeServiceImpl struct {
	contentLikeRepo repository.ContentLikeRepo
}

func NewContentLikeService(contentLikeRepo repository.ContentLikeRepo) ContentLikeService {
	return &ContentLikeServiceImpl{
		contentLikeRepo: contentLikeRepo,
	}
}

func (cl ContentLikeServiceImpl) Like(c context.Context, types string, content_id int32, user_id int32) error {
	var contentLike entity.ContentLikeHistory
	contentLike.UserId = user_id
	contentLike.Type = types
	contentLike.ContentId = content_id

	if err := cl.contentLikeRepo.Like(c, &contentLike); err != nil {
		return err
	}

	return nil
}
