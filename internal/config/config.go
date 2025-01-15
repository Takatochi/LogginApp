package config

import "LoggingApp/pkg/database"

type Config struct {
	BindAddr   string          `toml:"bind_addr"`
	Connection database.Config `toml:"connection_strings"`
}
