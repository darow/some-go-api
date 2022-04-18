package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type AuthLogRepository struct {
	store *Store
}

//LogAuthenticateAttempt Записываем в БД модель события аутентификации.
func (r *AuthLogRepository) LogAuthenticateAttempt(e *model.AuthenticationLog) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO authorization_events (user_id, event_id) VALUES ($1, $2) RETURNING created_time;",
		e.UserID,
		e.Event,
	).Scan(
		&e.Timestamp,
	); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}

	return nil
}

//FailedAttemptsCount Находим количество записей о неудачных аутентификациях пользователя в БД.
func (r *AuthLogRepository) FailedAttemptsCount(u *model.User) (count int, err error) {
	count = -1
	if u.ID == 0 {
		return count, errors.New("user ID must be not 0")
	}

	if err = r.store.db.QueryRow(
		"SELECT COUNT(*) FROM authorization_events WHERE user_id = $1 AND event_id = $2;",
		u.ID,
		model.AuthorizeWrongPassword,
	).Scan(
		&count,
	); err != nil {
		if err == sql.ErrNoRows {
			return count, store.ErrRecordNotFound
		}
		return count, err
	}

	return  count, nil
}

//GetAuthenticateHistory Формируем список записей об аутентификации пользователя из БД .
func (r *AuthLogRepository) GetAuthenticateHistory(u *model.User) (logs []*model.AuthenticationLog, err error) {
	rows, err := r.store.db.Query(
		"SELECT created_time, event_id FROM authorization_events WHERE user_id = $1;",
		u.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {
		log := &model.AuthenticationLog{}
		if err := rows.Scan(&log.Timestamp, &log.Event); err != nil {
			logrus.Warn(err)
			continue
		}

		logs = append(logs, log)
	}

	return logs, nil
}

//DeleteAuthorizeHistory Удаляем все записи об авторизации по id пользователя
func (r *AuthLogRepository) DeleteAuthorizeHistory(u *model.User) error {
	if err := r.store.db.QueryRow(
		"DELETE FROM authorization_events WHERE user_id = $1;",
		u.ID,
	).Scan(); err != sql.ErrNoRows {
		return err
	}
	return nil
}
