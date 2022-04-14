package sqlstore

import (
	"database/sql"
	"errors"
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

func (r *UserRepository) LogAuthenticateAttempt(e *model.AuthorizationEvent) error {
	if e.UserID == 0 {
		return errors.New("user ID must be not 0")
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO authorization_events (user_id, event) VALUES ($1, $2) RETURNING timestamp;",
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
		"SELECT COUNT(*) FROM authorization_events WHERE user_id = $1 AND event = $2;",
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

//func (r *UserRepository) CheckPass(u *model.User, pass string) (isHash bool, err error) {
//	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass))
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}

//func (r *UserRepository) FindByLoginPass(login, pass string) (u *model.User, err error) {
//	u, err = r.FindByLogin(login)
//	if err != nil {
//		return nil, err
//	}
//	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass))
//	if err != nil {
//		return nil, err
//	}
//	return u, nil
//}
