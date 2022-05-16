package schedule

import "shub_go/src/enums"

type CreateScheduleInput struct {
	Title     string             `form:"title" binding:"required"`
	StartTime string             `form:"startTime" binding:"required"`
	Day       string             `form:"day" binding:"required"`
	Link      string             `form:"link"`
	Type      enums.TypeSchedule `form:"type" binding:"required"`
	ClassId   int                `form:"classId" binding:"required"`
}
