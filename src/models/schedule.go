package models

import (
	"shub_go/src/enums"
	"time"
)

type Schedule struct {
	Id        int
	Title     string
	StartTime string
	Day       *time.Time
	Link      string
	Type      enums.TypeSchedule
	ClassId   int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
