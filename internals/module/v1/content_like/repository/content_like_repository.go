package repository

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"gorm.io/gorm"
)

type ContentLikeRepo interface {
	Like(c context.Context, payload *entity.ContentLikeHistory) error
	GetLikeById(c context.Context, content_id int32, user_id int32) (entity.ContentLikeHistory, error)
	Update(c context.Context, content_id int32, user_id int32, types string) error
	Delete(c context.Context, content_id int32, user_id int32) error
}

type ContentLikeRepoImpl struct {
	db *gorm.DB
}

func NewContentLikeRepo(db *gorm.DB) ContentLikeRepo {
	return &ContentLikeRepoImpl{
		db: db,
	}
}

func (cl ContentLikeRepoImpl) Like(c context.Context, payload *entity.ContentLikeHistory) error {
	result := cl.db.WithContext(c).Create(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (cl ContentLikeRepoImpl) GetLikeById(c context.Context, content_id int32, user_id int32) (entity.ContentLikeHistory, error) {
	var contentLike entity.ContentLikeHistory

	result := cl.db.WithContext(c).Where("content_id =?", content_id).Where("user_id =?", user_id).Find(&contentLike)

	if result.Error != nil {
		return entity.ContentLikeHistory{}, result.Error
	}

	return contentLike, nil

}

func (cl ContentLikeRepoImpl) Update(c context.Context, content_id int32, user_id int32, types string) error {
	var contentLike entity.ContentLikeHistory
	contentLike.Type = types

	result := cl.db.WithContext(c).Where("content_id =?", content_id).Where("user_id =?", user_id).Updates(&contentLike)

	if result.Error != nil {
		return errors.New("error when updating like status")
	}

	return nil
}

func (cl ContentLikeRepoImpl) Delete(c context.Context, content_id int32, user_id int32) error {
	var contentLike entity.ContentLikeHistory
	if err := cl.db.WithContext(c).Where("content_id =?", content_id).Where("user_id =?", user_id).Delete(&contentLike).Error; err != nil {
		return err
	}
	return nil
}
