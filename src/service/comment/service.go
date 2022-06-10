package comment

import (
	"context"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, userId int, input CreateCommentInput) (*models.Comment, error) {
	comment := models.Comment{
		Content: input.Content,
		PostId:  input.PostId,
		UserId:  userId,
	}

	err := s.repository.Create(ctx, &comment)

	if err != nil {
		return nil, app.ErrInternalServer(err)
	}

	pusher := config.Conf.GetPusher()

	pusher.Trigger("post-1", "create-comment", comment)

	return &comment, nil

}
