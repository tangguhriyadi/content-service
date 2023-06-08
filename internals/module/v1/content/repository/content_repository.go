package repository

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/dto"
	"gorm.io/gorm"
)

type ContentRepository interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentPaginate, error)
	Create(c context.Context, payload *dto.ContentCreate) error
	GetById(c context.Context, id int) (entity.Content, error)
	Update(c context.Context, id int, payload *dto.Content) error
	Delete(c context.Context, id int, payload *entity.Content) error
}

type ContentRepositoryImpl struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &ContentRepositoryImpl{
		db: db,
	}
}

func (ct ContentRepositoryImpl) GetAll(c context.Context, page int, limit int) (dto.ContentPaginate, error) {
	var contents []entity.Content
	var count int64

	countResult := ct.db.WithContext(c).Model(&[]entity.Content{}).Where("deleted =?", false).Count(&count)
	if countResult.Error != nil {
		return dto.ContentPaginate{}, nil
	}

	result := ct.db.WithContext(c).Select("id", "name", "like_count", "comment_count", "owner_id", "type_id", "is_premium").Where("deleted =?", false).Offset((page - 1) * limit).Limit(limit).Find(&contents)
	if result.Error != nil {
		return dto.ContentPaginate{}, nil
	}

	var response = dto.ContentPaginate{
		Data:       &contents,
		Page:       page,
		Limit:      limit,
		TotalItems: count,
	}
	return response, nil
}

func (ct ContentRepositoryImpl) Create(c context.Context, payload *dto.ContentCreate) error {
	var content = entity.Content{
		Name:      payload.Name,
		IsPremium: payload.IsPremium,
		OwnerID:   payload.OwnerID,
	}

	result := ct.db.WithContext(c).Create(&content)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ct ContentRepositoryImpl) GetById(c context.Context, id int) (entity.Content, error) {
	var content entity.Content

	result := ct.db.WithContext(c).Where("deleted =?", false).First(&content, id)

	if result.Error != nil {
		return entity.Content{}, result.Error
	}

	return content, nil
}

func (ct ContentRepositoryImpl) Update(c context.Context, id int, payload *dto.Content) error {

	result := ct.db.WithContext(c).Where("id =?", id).Updates(payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ct ContentRepositoryImpl) Delete(c context.Context, id int, payload *entity.Content) error {

	result := ct.db.WithContext(c).Where("id =?", id).Updates(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
