package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		logrus.Warnf("load config toml file failed | %s", err)
		logrus.Info("using default config")
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		logrus.Fatal(err)
	}
}