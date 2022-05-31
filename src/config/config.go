package config

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	errpkg "shub_go/src/packages/err"
)

type Config struct {
	dbUrl     string
	port      string
	db        *gorm.DB
	secretKey string
	sql       string
}

func (c *Config) GetSecretKey() string {
	return c.secretKey
}

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) GetDB() *gorm.DB {
	return c.db
}

var Conf Config

func Load() (*Config, error) {
	path, _ := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	sql := viper.GetString("DB_SQL")
	if sql != "postgres" && sql != "mysql" && sql != "sqlite" {
		return nil, errors.New("db sql not valid")
	}

	//db, err := gorm.Open(sqlite.Open(viper.GetString("DB_URL")), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	var db *gorm.DB

	if sql == "postgres" {
		db, err = gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if sql == "sqlite" {
		db, err = gorm.Open(sqlite.Open(viper.GetString("DB_URL")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if sql == "mysql" {
		db, err = gorm.Open(mysql.Open(viper.GetString("DB_URL")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		return nil, errors.New("db sql not valid")
	}

	if err != nil {
		return nil, err
	}

	result := &Config{
		dbUrl:     viper.GetString("DB_URL"),
		port:      viper.GetString("PORT"),
		db:        db,
		secretKey: viper.GetString("SECRET_KEY_JWT"),
	}

	Conf = *result

	errpkg.LoadError()

	return result, nil
}
