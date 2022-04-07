package teststore

import (
	"golang.org/x/crypto/bcrypt"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	r.users[u.Login] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u, ok := r.users[login]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (r *UserRepository) FindByLoginPass(login, pass string) (*model.User, error) {
	u, ok := r.users[login]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass)); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}