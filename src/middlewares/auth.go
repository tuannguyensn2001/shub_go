package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"shub_go/src/app"
	"shub_go/src/config"
	jwtpkg "shub_go/src/packages/jwt"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Authorization not valid")
	}

	return parts[1], nil
}

func Auth(ctx *gin.Context) {
	token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

	if err != nil {
		panic(app.ErrNoPermission(err))
		return
	}

	userId, err := jwtpkg.ValidateToken(config.Conf.GetSecretKey(), token)

	if err != nil {
		panic(app.ErrNoPermission(err))
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
