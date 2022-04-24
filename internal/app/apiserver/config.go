package apiserver

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

	p := path.Join(GetProjectRootPath(), configPath)

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


//GetProjectRootPath Получение пути корня проекта.
//Не будет работать, если этот файл переместить внутри проекта в другое место и не изменить layersCount
//Возможно стоит сделать отдельный пакет для этой функции, который нельзя будет перемещать
func GetProjectRootPath() string {
	const layersCountToRemove = 4
	_, p, _, _ := runtime.Caller(0)
	i := len(p) - 1
	j := layersCountToRemove

	for i >= 1 && j > 0 {
		if os.IsPathSeparator(p[i]) {
			j--
		}
		i--
	}
	dir := filepath.Clean(p[: i+1])
	return  dir
}


