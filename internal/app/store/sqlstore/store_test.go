package sqlstore_test

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"some-go-api/internal/app/apiserver"
	"strings"
	"testing"
)

var (
	psqlInfo string
)

type config struct {
	C          *apiserver.Config
	DBName     string `json:"db_name"`
	TestDBName string `json:"test_db_name"`
}

//TestMain функция запускается перед тестами sqlstore_test
func TestMain(m *testing.M) {
	//psqlInfo = os.Getenv("PSQL_INFO")

	config, err := newConfig("configs/apiserver.json")
	if err != nil {
		log.Warn(err)
	}

	psqlInfo = strings.Replace(config.C.PsqlInfo, config.DBName, config.TestDBName, 1)

	if psqlInfo == "" {
		log.Warn("internal/app/store/sqlstore/store_test.go can't read config DB file. Setting default values")
		psqlInfo = "host=localhost port=5432 user=postgres password=1 dbname=some_go_api_db_test sslmode=disable"
	}
	os.Exit(m.Run())
}

// newConfig Находим файл конфигурации по указанному пути. Создаем экземпляр конфикурации с параметрами из файла
//или значениями по умолчанию. Если файл не удалось найти или прочитать, то логируем причину в консоль.
func newConfig(configPath string) (*config, error) {
	c := &apiserver.Config{}

	conf := &config{
		c,
		"some_go_api_db_test",
		"db_name",
	}

	s, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	dir := filepath.Dir(s)

	// ¯\_(ツ)_/¯
	root := path.Base(filepath.Dir(path.Base(filepath.Dir(path.Base(filepath.Dir(path.Base(dir)))))))

	p := path.Join(root, configPath)

	confFile, err := os.Open(p)
	if err != nil {
		log.Warn("open config file failed| path:", p)
		return conf, err
	}

	byteValue, err := ioutil.ReadAll(confFile)
	if err != nil {
		log.Warn("ReadAll config file failed")
		return conf, err
	}

	if json.Unmarshal(byteValue, c); err != nil {
		log.Warn("unmarshal config c failed")
		return conf, err
	}

	if json.Unmarshal(byteValue, conf); err != nil {
		log.Warn("unmarshal config conf failed")
		return conf, err
	}

	return conf, nil
}
