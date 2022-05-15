package models

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID               int        `gorm:"column:id;" json:"id"`
	Name             string     `gorm:"column:name;" json:"name"`
	Code             string     `gorm:"column:code;" json:"code"`
	ApproveStudent   int        `gorm:"column:approve_student;" json:"approveStudent" `
	PreventQuitClass int        `gorm:"column:prevent_quit_class;" json:"preventQuitClass"`
	ShowMark         int        `gorm:"column:show_mark;" json:"showMark"`
	DisableNewsfeed  int        `gorm:"column:disable_newsfeed;" json:"disableNewsfeed"`
	SubjectId        int        `gorm:"column:subject_id;" json:"subjectId"`
	GradeId          int        `gorm:"column:grade_id;" json:"gradeId"`
	UserId           int        `gorm:"column:user_id;" json:"userId"`
	CreatedAt        *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt        *gorm.DeletedAt
	Grade            Grade   `json:"grade"`
	Subject          Subject `json:"subject"`
	PrivateCode      *string `gorm:"column:private_code;" json:"-" `
}
