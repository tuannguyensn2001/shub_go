package comment

import (
	"context"
	"gorm.io/gorm"
	"shub_go/src/models"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, comment *models.Comment) error {
	if err := r.db.Create(comment).Error; err != nil {
		return err
	}

	return nil
}
