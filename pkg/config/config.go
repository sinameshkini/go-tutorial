package config

import "go_tutorial/pkg/database"

type Config struct {
	Server   Server
	Database database.Config
}

type Server struct {
	Address string
}

//havent used it yet
type JWS struct {
	secret_key       string
	token_expiration int
}
