package models

import (
	"shub_go/src/enums"
	"time"
)

type Schedule struct {
	Id        int                `gorm:"column:id;" json:"id"`
	Title     string             `gorm:"column:title;" json:"title"`
	StartTime string             `gorm:"column:start_time;" json:"startTime"`
	Day       *time.Time         `gorm:"column:day;" json:"day"`
	Link      string             `gorm:"column:link;" json:"link"`
	Type      enums.TypeSchedule `gorm:"column:type;'" json:"type"`
	ClassId   int                `gorm:"column:class_id;" json:"classId"`
	CreatedAt *time.Time         `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time         `gorm:"column:updated_at" json:"updatedAt"`
}
