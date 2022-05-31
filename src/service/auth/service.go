package auth

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/models"
	hashpkg "shub_go/src/packages/hash"
	jwtpkg "shub_go/src/packages/jwt"
)

type service struct {
	repository IRepository
}

type IService interface {
	Register(ctx context.Context, input RegisterInput) (*models.User, error)
	Login(ctx context.Context, input LoginInput) (*LoginOutput, error)
	GetUserById(ctx context.Context, userId int64) (*models.User, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Register(ctx context.Context, input RegisterInput) (*models.User, error) {
	check, err := s.repository.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if check != nil {
		return nil, app.ErrEntityExisted(errors.New("User existed"), "User existed")
	}

	user, err := s.repository.Create(ctx, input)

	if err != nil {
		return nil, app.NewErrorResponse("Register failed", http.StatusInternalServerError, nil, err)
	}

	result, err := s.repository.FindById(ctx, user.Id)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *service) Login(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	user, err := s.repository.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	check := hashpkg.Compare(input.Password, user.Password)

	if !check {
		return nil, app.ErrInvalidRequestWithMessage(errors.New("Wrong password"), "Email or password not valid")
	}

	accessToken, err := jwtpkg.GenerateToken(config.Conf.GetSecretKey(), user.Id, 10000)

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwtpkg.GenerateToken(config.Conf.GetSecretKey(), user.Id, 20000)

	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

func (s *service) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	return s.repository.FindById(ctx, int(userId))
}
