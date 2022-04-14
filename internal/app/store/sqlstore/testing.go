package sqlstore

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestDB(t *testing.T, psqlInfo string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	log.Info("test db Ping Success")

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
