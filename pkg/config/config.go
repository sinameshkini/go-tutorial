package config

import "go_tutorial/pkg/database"

type Config struct {
	Server   Server
	Database database.Config
	Jws      JWS
}

type Server struct {
	Address string
}

type JWS struct {
	Secret_key       string
	Token_expiration int
}
