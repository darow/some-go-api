package apiserver

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Config ...
type Config struct {
	LogLevel string `json:"log_level"`
	BindAddr string `json:"bind_addr"`
	PsqlInfo string `json:"psql_info"`
}

// NewConfig Находим файл конфигурации по указанному пути. Создаем экземпляр конфикурации с параметрами из файла
//или значениями по умолчанию. Если файл не удалось найти или прочитать, то логируем причину в консоль.
func NewConfig(configPath string) (*Config, error) {
	conf := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}

	ex, err := os.Executable()
	if err != nil {
		log.Println(err)
	}

	exPath := filepath.Dir(ex)
	p := path.Join(exPath, configPath)

	confFile, err := os.Open(p)
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


