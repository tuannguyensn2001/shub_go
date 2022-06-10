package manage_class

import "shub_go/src/enums"

type CreateClassInput struct {
	Name             string `form:"name" binding:"required"`
	Code             string `form:"code" `
	ApproveStudent   bool   `form:"approveStudent" `
	PreventQuitClass bool   `form:"preventQuitClass" `
	ShowMark         bool   `form:"showMark" `
	DisableNewsfeed  bool   `form:"disableNewsfeed" `
	SubjectId        int    `form:"subjectId" `
	GradeId          int    `form:"gradeId" `
	PrivateCode      string `form:"privateCode"`
}

type QueryClass struct {
	name      *string
	direction enums.Direction
	orderBy   *string
}

type QueryClassOutput struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type AddMemberInput struct {
	UserId int `form:"userId" binding:"required"`
}
