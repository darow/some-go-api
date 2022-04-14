package teststore

import (
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
	authAttempts map[int]*model.AuthorizationLog
}

func (r *UserRepository) Create(u *model.User) error {
	u.BeforeCreate()
	u.ID = len(r.users)
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	for _, u := range r.users {
		if u.Login == login {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) LogAuthenticateAttempt(event *model.AuthorizationLog) error {
	id := len(r.authAttempts)
	r.authAttempts[id] = event
	return nil
}

func (r *UserRepository) FailedAttemptsCount(u *model.User) (count int, err error) {
	for _, v := range r.authAttempts {
		if v.UserID == u.ID {
			if v.Event == model.AuthorizeWrongPassword {
				count++
			}
		}
	}
	return  count, nil
}

//func (r *UserRepository) FindByLoginPass(login, pass string) (*model.User, error) {
//	for _, u := range r.users {
//		if u.Login == login {
//			if bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass)) == nil {
//				return u, nil
//			} else {
//				return nil, store.ErrRecordNotFound
//			}
//		}
//	}
//
//	return nil, store.ErrRecordNotFound
//}