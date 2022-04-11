package sqlstore

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	u.BeforeCreate()

	if err := r.store.db.QueryRow(
		"INSERT INTO users (login, encrypted_password, login_attempts) VALUES ($1, $2, $3) RETURNING user_id",
		u.Login,
		u.EncryptedPassword,
		0,
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

//func (r *UserRepository) CheckPass(u *model.User, pass string) (isHash bool, err error) {
//	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass))
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}

func (r *UserRepository) FindByLoginPass(login, pass string) (u *model.User, err error) {
	u, err = r.FindByLogin(login)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass))
	if err != nil {
		return nil, err
	}
	return u, nil
}
