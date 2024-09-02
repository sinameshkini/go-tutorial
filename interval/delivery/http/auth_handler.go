package http

import (
	"fmt"
	"gu_tutorial/interval/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	users []domain.User
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{users: []domain.User{}}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var newUser domain.User
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	for _, u := range h.users {
		if u.Email == newUser.Email {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
		}
	}

	h.users = append(h.users, newUser)
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

	var foundUser *domain.User
	for _, u := range h.users {
		if u.Email == credentials.Email {
			foundUser = &u
			break
		}
	}

	if foundUser == nil {
		return echo.ErrUnauthorized
	}

	if foundUser.Password != credentials.Password {
		return echo.ErrUnauthorized
	}

	token := fmt.Sprintf("mock-token-for-%s", credentials.Email)
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) ResetPassword(c echo.Context) error {
	var email string
	fmt.Println("hello")
	err := c.Bind(&email)
	fmt.Println(err)
	if err != nil {
		return echo.ErrBadRequest
	}

	found := false
	for _, u := range h.users {
		if u.Email == email {
			found = true
			break
		}
	}

	if !found {
		return echo.ErrNotFound
	}

	fmt.Printf("Reset password email sent to %s\n", email)
	return c.JSON(http.StatusOK, "Email sent with reset instructions")
}
