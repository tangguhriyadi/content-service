package service

import (
	"context"

	"github.com/tangguhriyadi/content-service/internals/module/v1/content/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/repository"
)

type ContentService interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentPaginate, error)
	Create(c context.Context, payload dto.ContentPayload, user_id int32) error
}

type contentServiceImpl struct {
	contentRepository repository.ContentRepository
}

func NewContentService(contentRepository repository.ContentRepository) ContentService {
	return &contentServiceImpl{
		contentRepository: contentRepository,
	}
}

func (cs contentServiceImpl) GetAll(c context.Context, page int, limit int) (dto.ContentPaginate, error) {
	result, err := cs.contentRepository.GetAll(c, page, limit)
	if err != nil {
		return dto.ContentPaginate{}, nil
	}
	return result, nil
}

func (cs contentServiceImpl) Create(c context.Context, payload dto.ContentPayload, user_id int32) error {
	var content dto.ContentCreate

	content.Name = payload.Name
	content.IsPremium = payload.IsPremium
	content.OwnerID = user_id

	err := cs.contentRepository.Create(c, &content)

	if err != nil {
		return err
	}

	return nil
}
