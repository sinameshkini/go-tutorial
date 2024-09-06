package main

import (
	handler "gu_tutorial/interval/delivery/http"
	"gu_tutorial/interval/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.StartDB()
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository(db)
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

	authHandler := handler.NewAuthHandler(*userRepository)

	e.POST("/auth/signup", authHandler.SignUp)
	e.POST("/auth/signin", authHandler.SignIn)
	e.POST("/auth/reset-password", authHandler.ResetPassword)

	e.Logger.Fatal(e.Start(":1323"))
}
