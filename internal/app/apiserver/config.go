package apiserver

// Config ...
type Config struct {
	BindAddr  string `toml:"bind_addr"`
	LogLevel  string `toml:"log_level"`
	Psql_info string `toml:"psql_info"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
