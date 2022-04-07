package sqlstore

type Config struct {
	psqlInfo string `toml:"psql_info"`
	//host     string
	//port     string
	//user     string
	//password string
	//dbname   string
}

func NewConfig() *Config {
	return &Config{
		psqlInfo: "host=localhost port=5432 user=postgres password=1 dbname=some_go_api_db sslmode=disable",
	}
}
