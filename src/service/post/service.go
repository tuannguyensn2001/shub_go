package post

import (
	"context"
	"shub_go/src/enums"
	"shub_go/src/models"
)

type IService interface {
	Create(ctx context.Context, input Input, userId int) (*models.Post, error)
	FindByClassId(ctx context.Context, classId int, query query) (*queryOutput, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, input Input, userId int) (*models.Post, error) {
	post := models.Post{
		Content: input.Content,
		ClassId: input.ClassId,
		UserId:  userId,
	}

	post.SetIsShow(enums.IsShow)

	err := s.repository.Create(ctx, &post)

	if err != nil {
		return nil, err
	}

	result, err := s.repository.FindById(ctx, post.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) FindByClassId(ctx context.Context, classId int, query query) (*queryOutput, error) {
	return s.repository.FindByClassId(ctx, classId, query)
}
