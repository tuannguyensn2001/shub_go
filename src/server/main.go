package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"shub_go/src/cmd"
	"shub_go/src/config"
	"shub_go/src/middlewares"
	"shub_go/src/routes"
)

func main() {
	_, err := config.Load()

	if err != nil {
		log.Fatalln("config err", err)
	}

	rootCmd := cmd.GetRoot()

	rootCmd.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {

			conf := config.Conf
			r := gin.Default()

			r.Use(middlewares.Recover)
			r.Use(middlewares.Cors)

			routes.Routes(r)
			err := r.Run(":" + conf.GetPort())

			defer func(mq *amqp.Connection) {
				err := mq.Close()
				if err != nil {

				}
			}(config.Conf.GetRabbitMq())

			if err != nil {
				return
			}
		},
	})

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln("err command", err)
	}

}

//package main
//
//import (
//	"context"
//	"flag"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"log"
//	postpb "shub_go/src/proto/post"
//)
//
//var (
//	addr = flag.String("addr", "localhost:2906", "the address post service")
//	name = flag.String("name", "post service", "hehe")
//)
//
//func main() {
//	flag.Parse()
//
//	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
//
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//
//	defer conn.Close()
//
//	c := postpb.NewPostServiceClient(conn)
//
//	r, err := c.Create(context.Background(), &postpb.CreatePostRequest{
//		Content: "hello ban oi",
//	})
//
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("Greeting: %s", r)
//}
