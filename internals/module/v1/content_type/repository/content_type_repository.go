package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/dto"
	"gorm.io/gorm"
)

type ContentTypeRepository interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentTypePaginate, error)
	GetById(c context.Context, id int) (entity.ContentType, error)
	Create(c context.Context, payload *dto.ContentTypePayload, user_id int32) error
	Update(c context.Context, id int, payload *dto.ContentTypePayload) error
	Delete(c context.Context, payload *entity.ContentType, id int) error
}

type ContentTypeRepositoryImpl struct {
	db *gorm.DB
}

func NewContentTypeRepository(db *gorm.DB) ContentTypeRepository {
	return &ContentTypeRepositoryImpl{
		db: db,
	}
}

func (ct ContentTypeRepositoryImpl) GetAll(c context.Context, page int, limit int) (dto.ContentTypePaginate, error) {
	var contentTypes []entity.ContentType
	var count int64

	countResult := ct.db.WithContext(c).Model(&[]entity.ContentType{}).Where("deleted =?", false).Count(&count)
	if countResult.Error != nil {
		return dto.ContentTypePaginate{}, nil
	}

	result := ct.db.WithContext(c).Select("id", "name").Where("deleted =?", false).Offset((page - 1) * limit).Limit(limit).Find(&contentTypes)
	if result.Error != nil {
		return dto.ContentTypePaginate{}, nil
	}

	var modifiedContentTypes []dto.ContentType

	for _, obj := range contentTypes {
		modifiedContentTypes = append(modifiedContentTypes, dto.ContentType{
			ID:   obj.ID,
			Name: obj.Name,
		})
	}

	var response = dto.ContentTypePaginate{
		Data:       &modifiedContentTypes,
		Page:       page,
		Limit:      limit,
		TotalItems: count,
	}

	return response, nil

}

func (ct ContentTypeRepositoryImpl) GetById(c context.Context, id int) (entity.ContentType, error) {
	var contentType entity.ContentType

	result := ct.db.WithContext(c).Where("deleted =?", false).First(&contentType, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entity.ContentType{}, nil
	}

	if result.Error != nil {
		return entity.ContentType{}, result.Error
	}

	return contentType, nil
}

func (ct ContentTypeRepositoryImpl) Update(c context.Context, id int, payload *dto.ContentTypePayload) error {
	var contentType entity.ContentType
	contentType.Name = payload.Name
	contentType.UpdatedAt = time.Now()
	result := ct.db.WithContext(c).Where("id =?", id).Updates(&contentType)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ct ContentTypeRepositoryImpl) Create(c context.Context, payload *dto.ContentTypePayload, user_id int32) error {
	var content = entity.ContentType{
		Name:      payload.Name,
		CreatedBy: user_id,
	}

	result := ct.db.WithContext(c).Create(&content)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ct ContentTypeRepositoryImpl) Delete(c context.Context, payload *entity.ContentType, id int) error {

	result := ct.db.WithContext(c).Where("id =?", id).Updates(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
