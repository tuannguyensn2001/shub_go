package schedule

import (
	"context"
	"gorm.io/gorm"
	"shub_go/src/models"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, schedule *models.Schedule) error
	FindByRange(ctx context.Context, classId int, start time.Time, end time.Time) ([]models.Schedule, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, schedule *models.Schedule) error {
	err := r.db.Create(schedule).Error

	return err

}

func (r *repository) FindByRange(ctx context.Context, classId int, start time.Time, end time.Time) ([]models.Schedule, error) {
	var result []models.Schedule

	err := r.db.Where("class_id = ?", classId).Where("day > ? ", start).Where("day < ?", end).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil

}
