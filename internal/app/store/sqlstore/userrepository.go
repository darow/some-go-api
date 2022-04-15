package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	u.BeforeCreate()

	if err := r.store.db.QueryRow(
		"INSERT INTO users (login, encrypted_password) VALUES ($1, $2) RETURNING user_id",
		u.Login,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT user_id, login, encrypted_password FROM users WHERE user_id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT user_id, login, encrypted_password FROM users WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) LogAuthenticateAttempt(e *model.AuthorizationLog) error {
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

func (r *UserRepository) FailedAttemptsCount(u *model.User) (count int, err error) {
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

func (r *UserRepository) GetAuthorizeHistory(u *model.User) (logs []*model.AuthorizationLog, err error) {
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

func (r *UserRepository) DeleteAuthorizeHistory(u *model.User) error {
	if err := r.store.db.QueryRow(
		"DELETE FROM authorization_events WHERE user_id = $1;",
		u.ID,
	).Scan(); err != nil {
		return err
	}
	return nil
}
