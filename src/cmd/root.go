package cmd

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"shub_go/src/config"
	"shub_go/src/models"
	"shub_go/src/packages/gen_code"
)

var cmdRoot = &cobra.Command{
	Use: "shub",
}

func GetRoot() *cobra.Command {
	cmdRoot.AddCommand(seedGrade())
	cmdRoot.AddCommand(seedSubject())
	cmdRoot.AddCommand(genErrorCode())
	return cmdRoot
}

func genErrorCode() *cobra.Command {
	return &cobra.Command{
		Use: "gen-error-code",
		Run: func(cmd *cobra.Command, args []string) {
			gen_code.GenCode()
		},
	}
}

func seedSubject() *cobra.Command {
	return &cobra.Command{
		Use: "seed-subject",
		Run: func(cmd *cobra.Command, args []string) {
			db := config.Conf.GetDB()

			err := db.Exec("DELETE FROM subjects").Error

			if err != nil {
				log.Fatalln("err", err)
			}

			subjects := []models.Subject{
				{
					Name: "Toán học",
				},
				{
					Name: "Văn học",
				},
				{
					Name: "Tiếng Anh",
				},
			}

			err = db.Create(&subjects).Error

			if err != nil {
				log.Fatalln("err", err)
			}

			log.Println("seed subject successfully")

		},
	}
}

func seedGrade() *cobra.Command {
	return &cobra.Command{
		Use: "seed-grade",
		Run: func(cmd *cobra.Command, args []string) {
			db := config.Conf.GetDB()

			err := db.Exec("DELETE FROM grades").Error

			if err != nil {
				log.Println("err", err)
				return
			}

			err = db.Transaction(func(tx *gorm.DB) error {
				grades := []models.Grade{
					{
						Name: "Lớp 6",
					},
					{
						Name: "Lớp 7",
					},
					{
						Name: "Lớp 8",
					},
				}

				err = db.Create(&grades).Error

				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				log.Println("err", err)
				return
			}

			log.Println("seed grade successfully")
		},
	}
}
