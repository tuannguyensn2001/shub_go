package auth

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserFromCtx(ctx context.Context) (int, error) {
	userId := ctx.Value("user_id")

	switch userId.(type) {
	case int:
		return userId.(int), nil
	default:
		return -1, errors.New("can't get user from ctx")
	}

}

func SetUserIntoCtx(ctx *gin.Context) (context.Context, error) {
	userId, ok := ctx.Get("user_id")
	if !ok {
		return nil, errors.New("user id not exist in context")
	}

	result := context.WithValue(ctx, "user_id", userId.(int))

	return result, nil
}

func GetUserFromGinContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get("user_id")

	if !ok {
		return -1, errors.New("user id not valid in ctx")
	}

	switch userId.(type) {
	case int:
		return userId.(int), nil
	default:
		return -1, errors.New("user id type not valid in ctx")
	}

}
