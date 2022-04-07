package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"some-go-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	//TODO Конфиг файл вместо toml поменять на json/yaml
	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Warnf("load config toml file failed | %s", err)
		log.Info("using default config")
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	} else {
		log.Info("server working now")
	}
}
