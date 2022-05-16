package schedule

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

func (t *transport) Create(ctx *gin.Context) {
	var input CreateScheduleInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	result, err := t.service.Create(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))

}
