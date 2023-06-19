package service

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/content-service/internals/entity"
	contentRepo "github.com/tangguhriyadi/content-service/internals/module/v1/content/repository"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_like/repository"
)

type ContentLikeService interface {
	Like(c context.Context, types string, content_id int32, user_id int32) error
}

type ContentLikeServiceImpl struct {
	contentLikeRepo repository.ContentLikeRepo
	contentRepo     contentRepo.ContentRepository
}

func NewContentLikeService(contentLikeRepo repository.ContentLikeRepo, contentRepo contentRepo.ContentRepository) ContentLikeService {
	return &ContentLikeServiceImpl{
		contentLikeRepo: contentLikeRepo,
		contentRepo:     contentRepo,
	}
}

func (cl ContentLikeServiceImpl) Like(c context.Context, types string, content_id int32, user_id int32) error {

	_, err := cl.contentRepo.GetById(c, int(content_id))
	if err != nil {
		return errors.New("content not found")
	}

	find, _ := cl.contentLikeRepo.GetLikeById(c, content_id, user_id)

	if find.ContentId == 0 {

		// return errors.New("not found")
		var contentLike entity.ContentLikeHistory
		contentLike.UserId = user_id
		contentLike.Type = types
		contentLike.ContentId = content_id

		err := cl.contentLikeRepo.Like(c, &contentLike)

		return err
	}

	if find.Type == types {
		if err := cl.contentLikeRepo.Delete(c, content_id, user_id); err != nil {
			return err
		}
	} else {
		if err := cl.contentLikeRepo.Update(c, content_id, user_id, types); err != nil {
			return err
		}
	}

	// if find.Type == "like" {
	// 	if err := cl.contentLikeRepo.Update(c, content_id, user_id, "dislike"); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	if err := cl.contentLikeRepo.Update(c, content_id, user_id, "like"); err != nil {
	// 		return err
	// 	}
	// }

	return nil

}
