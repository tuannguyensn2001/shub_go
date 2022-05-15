package auth

import (
	"context"
	"gorm.io/gorm"
	"shub_go/src/models"
	hashpkg "shub_go/src/packages/hash"
)

type IRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, data RegisterInput) (*models.User, error)
	FindById(ctx context.Context, id int) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = @email", map[string]interface{}{"email": email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Create(ctx context.Context, data RegisterInput) (*models.User, error) {

	password, err := hashpkg.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindById(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	query := r.db.Where("id = ?", id).First(&user)

	if err := query.Error; err != nil {
		return nil, err
	}

	return &user, nil

}
