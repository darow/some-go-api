package apiserver

import (
	"database/sql"
	"net/http"
	"some-go-api/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.Psql_info)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	s := newServer(store)
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
