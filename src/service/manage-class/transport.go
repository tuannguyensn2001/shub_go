package manage_class

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/enums"
	"shub_go/src/service/auth"
	"strconv"
)

type transport struct {
	service IService
}

func NewTransport() *transport {
	repository := NewRepository(config.Conf.GetDB())

	authRepository := auth.NewRepository(config.Conf.GetDB())
	authService := auth.NewService(authRepository)

	service := NewService(repository, authService)

	return &transport{service: service}
}

func (t *transport) GetOption(ctx *gin.Context) {
	option := ctx.Param("option")

	var result interface{}
	var err error

	if option == "subjects" {
		result, err = t.service.GetSubjects()
	} else if option == "grades" {
		result, err = t.service.GetGrades()
	} else {
		err = app.ErrInvalidRequest(errors.New("option not valid"))
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))

}

func (t *transport) CreateClass(ctx *gin.Context) {
	var input CreateClassInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	userId, _ := ctx.Get("user_id")

	result, err := t.service.CreateClass(ctx.Request.Context(), input, userId.(int))

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))
}

func (t *transport) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	result, err := t.service.GetById(ctx.Request.Context(), id)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", result))

}

func (t *transport) QueryByUserId(ctx *gin.Context) {
	params := QueryClass{}

	userId, _ := ctx.Get("user_id")

	if len(ctx.Query("orderBy")) > 0 {
		orderBy := ctx.Query("orderBy")
		params.orderBy = &orderBy
	}

	if len(ctx.Query("name")) > 0 {
		name := ctx.Query("name")
		params.name = &name
	}

	if len(ctx.Query("direction")) > 0 {
		direction := ctx.Query("direction")
		params.direction = enums.Direction(direction)
	} else {
		params.direction = enums.ASC
	}

	result, err := t.service.QueryByUserId(ctx.Request.Context(), userId.(int), params)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("success", QueryClassOutput{
		Total: len(result),
		Data:  result,
	}))

}

func (t *transport) AddStudentToClass(ctx *gin.Context) {
	classId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	var input AddMemberInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	teacherId, err := auth.GetUserFromGinContext(ctx)

	if err != nil {
		panic(err)
	}

	err = t.service.AddStudentToClass(ctx, input.UserId, classId, teacherId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("add successfully", nil))
}
