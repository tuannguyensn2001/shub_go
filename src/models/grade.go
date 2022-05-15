package models

import (
	"time"
)

type Grade struct {
	ID        int        `gorm:"column:id;" json:"id"`
	Name      string     `gorm:"column:name;" json:"name"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
