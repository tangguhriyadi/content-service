package service

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content_type/repository"
)

type ContentTypeService interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentTypePaginate, error)
	GetById(c context.Context, id int) (dto.ContentType, error)
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
