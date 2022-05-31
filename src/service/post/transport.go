package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shub_go/src/app"
	"shub_go/src/config"
	"shub_go/src/service/auth"
	"strconv"
)

type transport struct {
	service IService
}

func NewTransport() *transport {
	repository := NewRepository(config.Conf.GetDB())
	service := NewService(repository)

	return &transport{
		service: service,
	}
}

func (t *transport) Create(ctx *gin.Context) {
	var post Input

	if err := ctx.ShouldBind(&post); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	userId, err := auth.GetUserFromGinContext(ctx)

	if err != nil {
		panic(err)
	}

	result, err := t.service.Create(ctx, post, userId)

	output := Output{
		ID:        result.ID,
		Content:   result.Content,
		ClassId:   result.ClassId,
		UserId:    result.UserId,
		IsShow:    result.GetIsShow(),
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		User:      result.User,
		Class:     result.Class,
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Create post successfully", output))

}

func (t *transport) GetByClass(ctx *gin.Context) {
	classId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if page < 0 {
		page = 1
	}

	if limit > 20 {
		limit = 10
	}

	if len(ctx.Query("page")) == 0 && len(ctx.Query("limit")) == 0 {
		page = -1
		limit = -1
	}

	query := query{
		Page:  page,
		Limit: limit,
	}

	result, err := t.service.FindByClassId(ctx, classId, query)

	if err != nil {
		panic(err)
	}

	output := make([]Output, 0)

	for _, item := range result.Data {
		output = append(output, Output{
			ID:        item.ID,
			Content:   item.Content,
			User:      item.User,
			Class:     item.Class,
			IsShow:    item.GetIsShow(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			UserId:    item.UserId,
			ClassId:   item.ClassId,
		})
	}

	response := queryResponse{
		Total: result.Total,
		Data:  output,
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Get posts successfully", response))

}
