package models

import "time"

type Comment struct {
	Id        int        `gorm:"column:id;" json:"id"`
	Content   string     `gorm:"column:content;" json:"content"`
	UserId    int        `gorm:"column:user_id;" json:"userId"`
	PostId    int        `gorm:"column:post_id;" json:"postId"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}
