package main

import (
	handler "gu_tutorial/interval/delivery/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	authHandler := handler.NewAuthHandler()

	e.POST("/auth/signup", authHandler.SignUp)
	e.POST("/auth/signin", authHandler.SignIn)
	e.POST("/auth/reset-password", authHandler.ResetPassword)

	e.Logger.Fatal(e.Start(":1323"))
}
