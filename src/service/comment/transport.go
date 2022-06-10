package comment

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/models"
	"shub_go/src/service/auth"
)

type IService interface {
	Create(ctx context.Context, userId int, input CreateCommentInput) (*models.Comment, error)
}

type transport struct {
	service IService
}

func NewTransport() *transport {
	repository := NewRepository(config.Conf.GetDB())
	service := NewService(repository)
	transport := &transport{
		service: service,
	}
	return transport
}

func (t *transport) Create(ctx *gin.Context) {
	var input CreateCommentInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	userId, err := auth.GetUserFromGinContext(ctx)

	if err != nil {
		panic(err)
	}

	comment, err := t.service.Create(ctx, userId, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("create comment successfully", comment))

}
