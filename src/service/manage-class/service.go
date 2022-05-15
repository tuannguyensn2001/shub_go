package manage_class

import (
	"context"
	"errors"
	"shub_go/src/app"
	"shub_go/src/models"
	helperpkg "shub_go/src/packages/helper"
	strpkg "shub_go/src/packages/str"
	"strings"
)

type IService interface {
	GetSubjects() ([]models.Subject, error)
	GetGrades() ([]models.Grade, error)
	CreateClass(ctx context.Context, input CreateClassInput, userId int) (*models.Class, error)
	GetById(ctx context.Context, id int) (*models.Class, error)
	QueryByUserId(ctx context.Context, userId int, params QueryClass) ([]models.Class, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) GetSubjects() ([]models.Subject, error) {
	return s.repository.GetAllSubjects()
}

func (s *service) GetGrades() ([]models.Grade, error) {
	return s.repository.GetAllGrades()
}

func (s *service) CreateClass(ctx context.Context, input CreateClassInput, userId int) (*models.Class, error) {

	code := strings.ToUpper(strpkg.Random(5))

	checkCode, err := s.repository.FindByCode(ctx, code)

	if err != nil {
		return nil, err
	}

	if checkCode != nil {
		return nil, app.ErrConflict(errors.New("Class existed"), "class existed")
	}

	class := models.Class{
		Name:             input.Name,
		ApproveStudent:   helperpkg.ConvertBoolToInt(input.ApproveStudent),
		PreventQuitClass: helperpkg.ConvertBoolToInt(input.PreventQuitClass),
		ShowMark:         helperpkg.ConvertBoolToInt(input.ShowMark),
		DisableNewsfeed:  helperpkg.ConvertBoolToInt(input.DisableNewsfeed),
		SubjectId:        input.SubjectId,
		GradeId:          input.GradeId,
		UserId:           userId,
		Code:             code,
	}

	if len(input.PrivateCode) != 0 {
		class.PrivateCode = &input.PrivateCode
	}

	err = s.repository.CreateClass(ctx, &class)

	if err != nil {
		return nil, err
	}

	result, err := s.repository.FindByID(ctx, class.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetById(ctx context.Context, id int) (*models.Class, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) QueryByUserId(ctx context.Context, userId int, params QueryClass) ([]models.Class, error) {
	return s.repository.QueryByUserId(ctx, userId, params)
}
