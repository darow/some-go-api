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

func (r *AuthLogRepository) LogAuthenticateAttempt(e *model.AuthorizationLog) error {
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

func (r *AuthLogRepository) GetAuthorizeHistory(u *model.User) (logs []*model.AuthorizationLog, err error) {
	rows, err := r.store.db.Query(
		"SELECT created_time, event_id FROM authorization_events WHERE user_id = $1;",
		u.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {
		log := &model.AuthorizationLog{}
		if err := rows.Scan(&log.Timestamp, &log.Event); err != nil {
			logrus.Warn(err)
			continue
		}
		logrus.Info(log.Timestamp)
		logs = append(logs, log)
	}

	return logs, nil
}

func (r *AuthLogRepository) DeleteAuthorizeHistory(u *model.User) error {
	if err := r.store.db.QueryRow(
		"DELETE FROM authorization_events WHERE user_id = $1;",
		u.ID,
	).Scan(); err != nil {
		return err
	}
	return nil
}
