package middlewares

import (
	"github.com/gin-gonic/gin"
	"shub_go/src/app"
	errpkg "shub_go/src/packages/err"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Header("Content-Type", "application/json")

			if appErr, ok := err.(*app.Error); ok {
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return
			}

			if appErr, ok := err.(*errpkg.Error); ok {
				ctx.AbortWithStatusJSON(int(appErr.Status), appErr)
				return
			}

			appErr := app.ErrInternalServer(err.(error))
			ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)

			return
		}
	}()

	ctx.Next()
}
