package store

import "some-go-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByLogin(string) (*model.User, error)
	FindByLoginPass(string, string) (*model.User, error)
}