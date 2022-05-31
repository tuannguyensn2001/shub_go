package models

import "time"

type Profile struct {
	ID        int        `gorm:"column:id;" json:"id"`
	Name      string     `gorm:"column:name;" json:"name"`
	Avatar    string     `gorm:"column:avatar;" json:"avatar"`
	UserId    int        `gorm:"column:user_id;" json:"userId"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
