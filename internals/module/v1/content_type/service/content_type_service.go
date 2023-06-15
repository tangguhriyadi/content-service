package service

import (
	"context"
	"errors"

	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/repository"
)

type ContentTypeService interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentTypePaginate, error)
	GetById(c context.Context, id int) (dto.ContentType, error)
	Update(c context.Context, id int, payload *dto.ContentTypePayload) error
	Create(c context.Context, payload *dto.ContentTypePayload, user_id int32) error
}

type ContentTypeServiceImpl struct {
	contentTypeRepository repository.ContentTypeRepository
}

func NewContentTypeService(contentTypeRepository repository.ContentTypeRepository) ContentTypeService {
	return &ContentTypeServiceImpl{
		contentTypeRepository: contentTypeRepository,
	}
}

func (ct ContentTypeServiceImpl) GetAll(c context.Context, page int, limit int) (dto.ContentTypePaginate, error) {
	result, err := ct.contentTypeRepository.GetAll(c, page, limit)
	if err != nil {
		return dto.ContentTypePaginate{}, nil
	}
	return result, nil
}

func (ct ContentTypeServiceImpl) GetById(c context.Context, id int) (dto.ContentType, error) {
	result, err := ct.contentTypeRepository.GetById(c, id)
	if err != nil {
		return dto.ContentType{}, err
	}

	var contentType dto.ContentType

	contentType.ID = result.ID
	contentType.Name = result.Name

	return contentType, nil
}

func (ct ContentTypeServiceImpl) Create(c context.Context, payload *dto.ContentTypePayload, user_id int32) error {
	var content dto.ContentTypePayload

	content.Name = payload.Name

	err := ct.contentTypeRepository.Create(c, &content, user_id)

	if err != nil {
		return err
	}

	return nil
}

func (ct ContentTypeServiceImpl) Update(c context.Context, id int, payload *dto.ContentTypePayload) error {
	_, err := ct.contentTypeRepository.GetById(c, id)
	if err != nil {
		return errors.New("content type not found")
	}

	var contentType dto.ContentTypePayload
	contentType.Name = payload.Name

	errors := ct.contentTypeRepository.Update(c, id, &contentType)
	if errors != nil {
		return errors
	}

	return nil
}
