package sqlstore_test

import (
	"os"
	"testing"
)

var (
	psqlInfo string
)

func TestMain(m *testing.M) {
	psqlInfo = os.Getenv("PSQL_INFO")
	if psqlInfo == "" {
		psqlInfo = "host=localhost port=5432 user=postgres password=1 dbname=some_go_api_db_test sslmode=disable"
	}
	os.Exit(m.Run())
}
