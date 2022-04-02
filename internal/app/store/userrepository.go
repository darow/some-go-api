package store

import "some-go-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (login, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Login,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByLogin(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM users WHERE login = $1",
		 email,
		 ).Scan(
			&u.ID,
			&u.Login,
			&u.EncryptedPassword,
		); err != nil {
			return nil, err
		}
	return u, nil
}
