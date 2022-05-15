package auth

import "shub_go/src/models"

type RegisterInput struct {
	Username string `form:"username" json:"username" binding:"required" `
	Email    string `form:"email" json:"email" binding:"required" validate:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required" validate:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginOutput struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         models.User `json:"user"`
}
