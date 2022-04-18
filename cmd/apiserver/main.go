package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"some-go-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.json", "path to config file")
}

func main() {
	flag.Parse()

	config, err := apiserver.NewConfig(configPath)
	if err != nil {
		log.Warn(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
