package schedule

import (
	"context"
	"errors"
	"shub_go/src/app"
	"shub_go/src/enums"
	"shub_go/src/models"
	timepkg "shub_go/src/packages/time"
	manage_class "shub_go/src/service/manage-class"
)

type IService interface {
	Create(ctx context.Context, input CreateScheduleInput, userId int) (*models.Schedule, error)
	GetFromRange(ctx context.Context, classId int, start string, end string) ([]models.Schedule, error)
}

type service struct {
	repository   IRepository
	classService manage_class.IService
}

func NewService(repository IRepository, classService manage_class.IService) *service {
	return &service{repository: repository, classService: classService}
}

func (s *service) Create(ctx context.Context, input CreateScheduleInput, userId int) (*models.Schedule, error) {

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

	class, err := s.classService.GetById(ctx, input.ClassId)

	if err != nil {
		return nil, err
	}

	if class.UserId != userId {
		return nil, app.ErrNoPermission(errors.New("user not valid"))
	}

	schedule := models.Schedule{
		Day:       t1,
		Title:     input.Title,
		Link:      input.Link,
		StartTime: input.StartTime,
		ClassId:   input.ClassId,
		Type:      input.Type,
	}

	err = s.repository.Create(ctx, &schedule)

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (s *service) GetFromRange(ctx context.Context, classId int, start string, end string) ([]models.Schedule, error) {

	startTime, err := timepkg.ParseDate(start)

	if err != nil {
		return nil, err
	}

	endTime, err := timepkg.ParseDate(end)

	if err != nil {
		return nil, err
	}

	result, err := s.repository.FindByRange(ctx, classId, *startTime, *endTime)

	if err != nil {
		return nil, err
	}

	return result, nil
}
