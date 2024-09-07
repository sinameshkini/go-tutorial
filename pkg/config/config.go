package config

import "go_tutorial/pkg/database"

type Config struct {
	Server   Server
	Database database.Config
}

type Server struct {
	Address string
}
