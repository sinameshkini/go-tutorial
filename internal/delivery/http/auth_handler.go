package http

import (
	"go_tutorial/internal/domain"
	"go_tutorial/internal/usecase"
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
	api.POST("/auth/signup", authHandler.SignUp)
	//api.POST("/auth/signin", authHandler.SignIn)
	//api.POST("/auth/reset-password", authHandler.ResetPassword)
}

func (h *handler) SignUp(c echo.Context) error {
	var newUser domain.User
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.usecase.SignUp(newUser.Email, newUser.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "User successfully registered")
}

//
//func (h *handler) SignIn(c echo.Context) error {
//
//	var credentials struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	if err := c.Bind(&credentials); err != nil {
//		return echo.ErrBadRequest
//	}
//
//	user, err := h.userRepository.FindUser(credentials.Email)
//	if err != nil {
//		return err
//	}
//
//	if user == nil || user.Password != credentials.Password {
//		return echo.ErrUnauthorized
//	}
//
//	token := fmt.Sprintf("mock-token-for-%s", credentials.Email)
//
//	return c.JSON(http.StatusOK, map[string]string{"token": token})
//}
//
//func (h *handler) ResetPassword(c echo.Context) error {
//	var email struct {
//		Email string `json:"email"`
//	}
//
//	if err := c.Bind(&email); err != nil {
//		return echo.ErrBadRequest
//	}
//
//	user, err := h.userRepository.FindUser(email.Email)
//	if err != nil {
//		return err
//	}
//
//	if user == nil {
//		return echo.ErrNotFound
//	}
//
//	return c.JSON(http.StatusOK, "Email sent with reset instructions")
//}
