package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
