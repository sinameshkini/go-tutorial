package main

import (
	"fmt"
	"go_tutorial/internal/delivery/http"
	handler "go_tutorial/internal/delivery/http"
	"go_tutorial/internal/repository"
	"go_tutorial/internal/usecase"
	"go_tutorial/pkg/config"
	"go_tutorial/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/golang-jwt/jwt/v5"
)

func main() {

	// Loading configs
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

	//Initializing database
	db, err := database.New(conf.Database)
	if err != nil {
		logrus.Fatalln("error connecting to database", err.Error())
	}
	if err = database.Migrate(db); err != nil {
		logrus.Errorln("error migrating database", err.Error())
	}

	// Initializing db and assigning db
	repo := repository.NewUserRepository(db)
	uc := usecase.New(repo)

	// Initializing the echo object
	e := echo.New()

	//Initializing addresses
	r := e.Group("/shop")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(http.JwtCustomClaims)
		},
		SigningKey: []byte(conf.Jws.Secret_key),
	}
	r.Use(echojwt.WithConfig(config))

	g := e.Group("/api/v1")

	//Initializing the handlers
	handler.Init(g, r, uc, &conf)

	//Starting the Server
	logrus.Fatalln(e.Start(conf.Server.Address))
}
