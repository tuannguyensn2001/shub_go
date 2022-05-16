package schedule

import (
	"context"
	"errors"
	"shub_go/src/app"
	"shub_go/src/enums"
	"shub_go/src/models"
	timepkg "shub_go/src/packages/time"
)

type IService interface {
	Create(ctx context.Context, input CreateScheduleInput) (*models.Schedule, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input CreateScheduleInput) (*models.Schedule, error) {

	if !enums.IsValidTypeSchedule(input.Type) {
		return nil, app.ErrInvalidRequest(errors.New("type not valid"))
	}

	if !timepkg.CheckValidHour(input.StartTime) {
		return nil, app.ErrInvalidRequest(errors.New("invalid start_time"))
	}

	t1, err := timepkg.ParseDate(input.Day)

	if err != nil {
		return nil, app.ErrInvalidRequest(errors.New("day not valid format"))
	}

	schedule := models.Schedule{
		Day: t1,
	}

	return &schedule, nil
}
