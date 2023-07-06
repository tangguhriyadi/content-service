package service

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/repository"
)

type ContentCommentService interface {
	GetByContentId(c context.Context, content_id int32) (*[]dto.ContentComment, error)
	PostComment(c context.Context, content_id int32, payload *dto.CommentPayload, user_id int32) error
}

type ContentCommentServiceImpl struct {
	contentCommentRepo repository.ContentCommentRepo
}

func NewContentCommentService(contentCommentRepo repository.ContentCommentRepo) ContentCommentService {
	return &ContentCommentServiceImpl{
		contentCommentRepo: contentCommentRepo,
	}
}

func (cc ContentCommentServiceImpl) GetByContentId(c context.Context, content_id int32) (*[]dto.ContentComment, error) {
	result, err := cc.contentCommentRepo.GetByContentId(c, content_id)
	if err != nil {
		return nil, err
	}

	resultlength := *result

	var contentComment []dto.ContentComment

	for _, v := range resultlength {
		reply, _ := cc.contentCommentRepo.GetReplies(c, v.ID)
		comment := dto.ContentComment{
			ID:        v.ID,
			UserId:    v.UserId,
			LikeCount: v.LikeCount,
			Comment:   v.Comment,
			ContentId: v.ContentId,
			Replies:   reply,
		}
		contentComment = append(contentComment, comment)
	}

	return &contentComment, nil
}

func (cc ContentCommentServiceImpl) PostComment(c context.Context, content_id int32, payload *dto.CommentPayload, user_id int32) error {
	var content entity.ContentComment

	content.Comment = payload.Comment
	content.ContentId = content_id
	content.UserId = user_id

	if err := cc.contentCommentRepo.PostComment(c, &content); err != nil {
		return err
	}

	return nil
}
