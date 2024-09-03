package main

import (
	handler "gu_tutorial/interval/delivery/http"
	"gu_tutorial/interval/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := repository.StartDB()
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository(db)
	e := echo.New()
	authHandler := handler.NewAuthHandler(*userRepository)

	e.POST("/auth/signup", authHandler.SignUp)
	e.POST("/auth/signin", authHandler.SignIn)
	e.POST("/auth/reset-password", authHandler.ResetPassword)

	e.Logger.Fatal(e.Start(":1323"))
}
