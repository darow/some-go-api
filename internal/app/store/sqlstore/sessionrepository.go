package sqlstore

import (
	"database/sql"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type SessionRepository struct {
	store *Store
}

func (r *SessionRepository) Create(s *model.Session) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO sessions (user_id, token, expiration_time) VALUES ($1, $2, $3) RETURNING session_id",
		s.UserID,
		s.Token,
		s.ExpirationTime,
	).Scan(&s.ID); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Find(token string) (*model.Session, error) {
	s := &model.Session{}
	if err := r.store.db.QueryRow(
		"SELECT session_id, user_id, token, expiration_time FROM sessions WHERE token = $1",
		token,
	).Scan(
		&s.ID,
		&s.UserID,
		&s.Token,
		&s.ExpirationTime,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

