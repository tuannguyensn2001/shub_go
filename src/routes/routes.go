package routes

import (
	"github.com/gin-gonic/gin"
	"shub_go/src/middlewares"
	"shub_go/src/service/auth"
	"shub_go/src/service/manage-class"
	"shub_go/src/service/schedule"
)

func Routes(r *gin.Engine) {

	authTransport := auth.NewTransport()
	classTransport := manage_class.NewTransport()
	scheduleTransport := schedule.NewTransport()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authTransport.Register)
		v1.POST("/auth/login", authTransport.Login)
		v1.GET("/auth/me", middlewares.Auth, authTransport.Me)

		v1.GET("/classes/option/:option", classTransport.GetOption)
		v1.POST("/classes", middlewares.Auth, classTransport.CreateClass)
		v1.GET("/classes/owner", middlewares.Auth, classTransport.QueryByUserId)
		v1.GET("/classes/:id", classTransport.GetById)

		v1.POST("/schedules", middlewares.Auth, scheduleTransport.Create)

	}
}
