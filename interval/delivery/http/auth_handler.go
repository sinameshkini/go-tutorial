package http

import (
	"fmt"
	"gu_tutorial/interval/domain"
	"gu_tutorial/interval/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userRepository repository.UserRepository
}

func NewAuthHandler(userRepository repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepository: userRepository}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var newUser domain.UserModel
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	usr, err := h.userRepository.FindUser(newUser.Email)
	if err != nil {
		return err
	}
	if usr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	err = h.userRepository.AddUser(&newUser)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "User successfully registered")
}

func (h *AuthHandler) SignIn(c echo.Context) error {

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.userRepository.FindUser(credentials.Email)
	if err != nil {
		return err
	}

	if user == nil || user.Password != credentials.Password {
		return echo.ErrUnauthorized
	}

	token := fmt.Sprintf("mock-token-for-%s", credentials.Email)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) ResetPassword(c echo.Context) error {
	var email struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&email); err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.userRepository.FindUser(email.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, "Email sent with reset instructions")
}
