package schedule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/service/auth"
	manage_class "shub_go/src/service/manage-class"
	"strconv"
)

type transport struct {
	service IService
}

func NewTransport() *transport {
	classRepository := manage_class.NewRepository(config.Conf.GetDB())

	authRepository := auth.NewRepository(config.Conf.GetDB())
	authService := auth.NewService(authRepository)

	classService := manage_class.NewService(classRepository, authService)

	repository := NewRepository(config.Conf.GetDB())
	service := NewService(repository, classService)

	return &transport{service: service}
}

func (t *transport) Create(ctx *gin.Context) {
	var input CreateScheduleInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	userId, _ := ctx.Get("user_id")

	result, err := t.service.Create(ctx.Request.Context(), input, userId.(int))

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))

}

func (t *transport) GetByRange(ctx *gin.Context) {
	classId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	start := ctx.Query("start")
	end := ctx.Query("end")

	result, err := t.service.GetFromRange(ctx.Request.Context(), classId, start, end)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))

}
