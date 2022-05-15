package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
)

type transport struct {
	service IService
}

func NewTransport() *transport {
	repository := NewRepository(config.Conf.GetDB())
	service := NewService(repository)

	return &transport{service: service}
}

func (t *transport) Register(ctx *gin.Context) {
	var input RegisterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	user, err := t.service.Register(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Register successfully", user))
}

func (t *transport) Login(ctx *gin.Context) {
	var input LoginInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	resp, err := t.service.Login(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Login successfully", resp))
}

func (t *transport) Me(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	user, err := t.service.GetUserById(ctx.Request.Context(), int64(userId.(int)))

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Get information user successfully", user))

}
