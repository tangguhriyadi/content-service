package repository

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_comment/dto"
	"gorm.io/gorm"
)

type ContentCommentRepo interface {
	GetByContentId(c context.Context, content_id int32) (*[]entity.ContentComment, error)
	GetReplies(c context.Context, content_id int32) (*[]dto.ContentCommentReply, error)
	PostComment(c context.Context, payload *entity.ContentComment) error
}

type ContentCommentImpl struct {
	db *gorm.DB
}

func NewContentCommentRepo(db *gorm.DB) ContentCommentRepo {
	return &ContentCommentImpl{
		db: db,
	}
}

func (cm ContentCommentImpl) GetByContentId(c context.Context, content_id int32) (*[]entity.ContentComment, error) {
	var contentComment []entity.ContentComment

	result := cm.db.WithContext(c).Where("content_id", content_id).Where("reply_to IS NULL").Find(&contentComment)

	if result.Error != nil {
		return nil, result.Error
	}

	return &contentComment, nil
}
func (cm ContentCommentImpl) GetReplies(c context.Context, comment_id int32) (*[]dto.ContentCommentReply, error) {
	var contentComment []entity.ContentComment

	if err := cm.db.WithContext(c).Where("reply_to =?", comment_id).Find(&contentComment).Error; err != nil {
		return nil, err
	}

	var replies []dto.ContentCommentReply

	for _, v := range contentComment {
		reply := dto.ContentCommentReply{
			ID:        v.ID,
			ContentId: v.ContentId,
			UserId:    v.UserId,
			LikeCount: v.LikeCount,
			Comment:   v.Comment,
			ReplyTo:   v.ReplyTo,
		}
		replies = append(replies, reply)
	}

	return &replies, nil
}

func (cm ContentCommentImpl) PostComment(c context.Context, payload *entity.ContentComment) error {
	if err := cm.db.WithContext(c).Create(&payload).Error; err != nil {
		return err
	}

	return nil
}
