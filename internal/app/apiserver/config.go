package apiserver

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

// Config ...
type Config struct {
	LogLevel string `json:"log_level"`
	BindAddr string `json:"bind_addr"`
	PsqlInfo string `json:"psql_info"`
}

// NewConfig ...
func NewConfig(configPath string) (*Config, error) {
	conf := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	confFile, err := os.Open(path.Join(wd, configPath))
	if err != nil {
		log.Warn("open config file failed")
		return conf, err
	}

	byteValue, err := ioutil.ReadAll(confFile)
	if err != nil {
		log.Warn("ReadAll config file failed")
		return conf, err
	}

	if json.Unmarshal(byteValue, conf); err != nil {
		log.Warn("unmarshal config failed")
		return conf, err
	}

	return conf, nil
}


