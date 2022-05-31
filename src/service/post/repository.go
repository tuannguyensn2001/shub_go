package post

import (
	"context"
	"gorm.io/gorm"
	"shub_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, post *models.Post) error
	FindById(ctx context.Context, id int) (*models.Post, error)
	FindByClassId(ctx context.Context, classId int, query query) (*queryOutput, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, post *models.Post) error {
	err := r.db.Create(post).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindById(ctx context.Context, id int) (*models.Post, error) {
	var post models.Post

	err := r.db.Preload("Class").Preload("User").Where("id = ?", id).First(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repository) FindByClassId(ctx context.Context, classId int, query query) (*queryOutput, error) {
	result := queryOutput{}

	var posts []models.Post

	offset := (query.Page - 1) * query.Limit

	sql := r.db.Preload("Class").Preload("User").Where("class_id = ?", classId)

	if query.Page == -1 && query.Limit == -1 {
		sql = sql.Find(&posts)
	} else {
		sql = sql.Offset(offset).Limit(query.Limit).Find(&posts)
	}

	err := sql.Error

	var total int64

	err = r.db.Table("posts").Count(&total).Error

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result.Data = posts
	result.Total = int(total)

	return &result, nil

}
