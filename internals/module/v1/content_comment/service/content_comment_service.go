package service

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/repository"
)

type ContentCommentService interface {
	GetByContentId(c context.Context, content_id int32) (*[]dto.ContentComment, error)
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
