package post

import (
	"shub_go/src/models"
	"time"
)

type UserOutput struct {
	Username string `json:"username"`
}

type Output struct {
	ID        int           `json:"id"`
	Content   string        `json:"content"`
	UserId    int           `json:"userId"`
	ClassId   int           `json:"classId"`
	IsShow    bool          `json:"isShow"`
	CreatedAt *time.Time    `json:"createdAt"`
	UpdatedAt *time.Time    `json:"updatedAt"`
	User      *models.User  `json:"user"`
	Class     *models.Class `json:"class"`
}

type Input struct {
	Content string `form:"content" binding:"required"`
	ClassId int    `form:"classId" binding:"required"`
}

type query struct {
	Limit int
	Page  int
}

type queryOutput struct {
	Total int
	Data  []models.Post
}

type queryResponse struct {
	Total int      `json:"total"`
	Data  []Output `json:"data"`
}
