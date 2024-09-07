package main

import (
	handler "go_tutorial/internal/delivery/http"
	"go_tutorial/internal/repository"
	"go_tutorial/internal/usecase"
	"go_tutorial/pkg/config"
	"go_tutorial/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// TODO read config from configs/config.yaml by Viper
	conf := config.Config{
		Server: config.Server{
			Address: ":8080",
		},
		Database: database.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "user",
			Password: "password",
			DBName:   "go_tutorial",
			Debug:    true,
		},
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

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	g := e.Group("/api/v1")

	handler.Init(g, uc)

	logrus.Fatalln(e.Start(conf.Server.Address))
}
