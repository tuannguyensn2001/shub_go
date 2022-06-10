package models

import "time"

type UserClass struct {
	UserId    int        `gorm:"column:user_id;" json:"userId"`
	ClassId   int        `gorm:"column:class_id;" json:"classId"`
	Role      int        `gorm:"column:role;" json:"role"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
