package http

import (
	"fmt"
	"go_tutorial/internal/domain"
	"go_tutorial/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase *usecase.AuthUsecase
}

func Init(g *echo.Group, uc *usecase.AuthUsecase) {
	authHandler := handler{
		usecase: uc,
	}

	api := g.Group("/auth")
	api.POST("/signup", authHandler.SignUp)
	api.POST("/signin", authHandler.SignIn)
	api.POST("/reset-password", authHandler.ResetPassword)
}

func (h *handler) SignUp(c echo.Context) error {
	var newUser domain.User
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}
	log.Print("help me")
	if err := h.usecase.SignUp(newUser.Email, newUser.Password); err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, "User successfully registered")
}

func (h *handler) SignIn(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.usecase.SignIn(user.Email, user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	token := fmt.Sprintf("mock-token-for-%s", user.Email)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *handler) ResetPassword(c echo.Context) error {

	var email domain.Email
	if err := c.Bind(&email); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.usecase.ResetPassowrd(email.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Email sent with reset instructions")
}
