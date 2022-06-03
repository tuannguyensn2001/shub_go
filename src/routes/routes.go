package routes

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"shub_go/src/middlewares"
	classpb "shub_go/src/proto/class"
	"shub_go/src/service/auth"
	"shub_go/src/service/manage-class"
	"shub_go/src/service/post"
	"shub_go/src/service/schedule"
)

var (
	port = flag.Int("port", 11000, "The server port")
)

func Routes(r *gin.Engine) {

	authTransport := auth.NewTransport()
	classTransport := manage_class.NewTransport()
	scheduleTransport := schedule.NewTransport()
	postTransport := post.NewTransport()

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

		if err != nil {
			log.Fatalf("err grpc tcp")
		}

		s := grpc.NewServer()

		classpb.RegisterClassServiceServer(s, classTransport)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

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
		v1.GET("/schedules/class/:id", scheduleTransport.GetByRange)

		v1.POST("/posts", middlewares.Auth, postTransport.Create)
		v1.GET("/posts/class/:id", postTransport.GetByClass)

	}
}
