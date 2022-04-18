package sqlstore_test

import (
	"github.com/sirupsen/logrus"
	"os"
	"some-go-api/internal/app/apiserver"
	"testing"
)

var (
	psqlInfo string
)

//TestMain функция запускается перед тестами sqlstore_test
func TestMain(m *testing.M) {
	//psqlInfo = os.Getenv("PSQL_INFO")
	logrus.Info("TestMain")
	config, err := apiserver.NewConfig("configs/apiserver.json")
	if err != nil {
		logrus.Warn(err)
	}

	psqlInfo = config.PsqlInfo

	if psqlInfo == "" {
		logrus.Warn("internal/app/store/sqlstore/store_test.go can't read config DB file. Setting default values")
		psqlInfo = "host=localhost port=5432 user=postgres password=1 dbname=some_go_api_db_test sslmode=disable"
	}
	os.Exit(m.Run())
}
