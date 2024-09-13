package http

import (
	"fmt"
	"go_tutorial/internal/domain"
	"go_tutorial/internal/usecase"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase *usecase.AuthUsecase
}

func Init(g *echo.Group, g2 *echo.Group, uc *usecase.AuthUsecase) {
	authHandler := handler{
		usecase: uc,
	}

	api := g.Group("/auth")
	api.POST("/signup", authHandler.SignUp)
	api.POST("/signin", authHandler.SignIn)
	api.POST("/reset-password", authHandler.ResetPassword)

	api2 := g2.Group("/restricted")

	api2.GET("/main", authHandler.Restricted)
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h *handler) SignUp(c echo.Context) error {
	// loads the new user data
	var newUser domain.User
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	//trying to signup the user
	if err := h.usecase.SignUp(newUser.Email, newUser.Password); err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, "User successfully registered")
}

func (h *handler) SignIn(c echo.Context) error {
	// loads the new user data
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}

	//trying to signin the user
	if err := h.usecase.SignIn(user.Email, user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//if succesful, we create a jwt token
	claims := &JwtCustomClaims{
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signing using the "secret in config"
	//plan: load the secret directly from config fie
	tk, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tk})
}

func (h *handler) ResetPassword(c echo.Context) error {
	//trying to load the email
	var email domain.Email
	if err := c.Bind(&email); err != nil {
		return echo.ErrBadRequest
	}

	//reset the old password with sth random
	var newPass, err = h.usecase.ResetPassowrd(email.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("your new password is : %s", newPass))
}

// restricted part that only loggedIn users should have access
func (h *handler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Email
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
