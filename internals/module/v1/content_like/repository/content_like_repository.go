package repository

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"gorm.io/gorm"
)

type ContentLikeRepo interface {
	Like(c context.Context, payload *entity.ContentLikeHistory) error
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
