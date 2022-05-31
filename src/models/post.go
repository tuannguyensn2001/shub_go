package models

import (
	"shub_go/src/enums"
	"time"
)

type Post struct {
	ID        int        `gorm:"column:id;" json:"id"`
	Content   string     `gorm:"column:content;" json:"content"`
	UserId    int        `gorm:"column:user_id;" json:"userId"`
	ClassId   int        `gorm:"column:class_id;" json:"classId"`
	isShow    enums.Show `gorm:"column:is_show;"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updatedAt"`
	User      *User
	Class     *Class
}

func (p *Post) GetIsShow() bool {
	return p.isShow == enums.IsShow
}

func (p *Post) SetIsShow(isShow enums.Show) {

	p.isShow = isShow
}
