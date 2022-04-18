package apiserver

import (
	"database/sql"
	"net/http"
	"some-go-api/internal/app/store/sqlstore"
)

// Start Получаем базу данных, передаем ее инициализатору sqlstore. Запускаем сервер, используя store. В
//случае ошибки запуска возвращаем эту ошибку для обработки.
func Start(config *Config) error {
	db, err := newDB(config.PsqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	s := newServer(store)
	return http.ListenAndServe(config.BindAddr, s)
}

// newDB Функция открытия базы данных. Берет строку конфигурации psqlInfo, открывает базу данных и проверяет
//успешность открытия с помощью db.Ping(). В случае успеха возвращает бд
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
