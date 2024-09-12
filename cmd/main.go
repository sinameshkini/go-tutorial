package main

import (
	"fmt"
	handler "go_tutorial/internal/delivery/http"
	"go_tutorial/internal/repository"
	"go_tutorial/internal/usecase"
	"go_tutorial/pkg/config"
	"go_tutorial/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	var conf config.Config
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
		return
	}

	db, err := database.New(conf.Database)
	if err != nil {
		logrus.Fatalln("error connecting to database", err.Error())
	}

	if err = database.Migrate(db); err != nil {
		logrus.Errorln("error migrating database", err.Error())
	}

	repo := repository.NewUserRepository(db)

	uc := usecase.New(repo)

	e := echo.New()

	g := e.Group("/api/v1")

	handler.Init(g, uc)

	logrus.Fatalln(e.Start(conf.Server.Address))
}
