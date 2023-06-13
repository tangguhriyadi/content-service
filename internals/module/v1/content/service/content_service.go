package service

import (
	"context"
	"errors"
	"time"

	"github.com/tangguhriyadi/content-service/internals/infrastructure/grpc"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/dto"
	"github.com/tangguhriyadi/content-service/internals/module/v1/content/repository"
)

type ContentService interface {
	GetAll(c context.Context, page int, limit int) (dto.ContentPaginate, error)
	Create(c context.Context, payload dto.ContentPayload, user_id int32) error
	GetById(c context.Context, id int) (dto.NewContent, error)
	Update(c context.Context, id int, payload *dto.ContentPayload) error
	Delete(c context.Context, id int) error
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

func (cs contentServiceImpl) GetById(c context.Context, id int) (dto.NewContent, error) {
	result, err := cs.contentRepository.GetById(c, id)

	if err != nil {
		return dto.NewContent{}, errors.New("content not found")
	}

	userGrpc, err := grpc.NewGrpcDial().GetUserData(result.OwnerID)
	if err != nil {
		return dto.NewContent{}, err
	}

	var user dto.User
	user.Age = userGrpc.Age
	user.Email = userGrpc.Email
	user.FullName = userGrpc.FullName

	var content dto.NewContent
	content.ID = result.ID
	content.CommentCount = result.CommentCount
	content.IsPremium = result.IsPremium
	content.LikeCount = result.LikeCount
	content.Name = result.Name
	content.Owner = &user
	content.TypeID = result.TypeID

	return content, nil
}

func (cs contentServiceImpl) Update(c context.Context, id int, payload *dto.ContentPayload) error {
	_, err := cs.contentRepository.GetById(c, id)
	if err != nil {
		return errors.New("content not found")
	}

	var content dto.Content

	content.IsPremium = payload.IsPremium
	content.Name = payload.Name

	errors := cs.contentRepository.Update(c, id, &content)
	if errors != nil {
		return errors
	}

	return nil
}

func (cs contentServiceImpl) Delete(c context.Context, id int) error {
	result, err := cs.contentRepository.GetById(c, id)

	if err != nil {
		return errors.New("content not found")
	}

	var now = time.Now()
	result.Deleted = true
	result.DeletedAt = &now

	if err := cs.contentRepository.Delete(c, id, &result); err != nil {
		return err
	}

	return nil
}
