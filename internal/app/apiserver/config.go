package apiserver

import "some-go-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string        `toml:"bind_addr"`
	LogLevel string        `toml:"log_level"`
	Store    *store.Config `toml:"store"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
