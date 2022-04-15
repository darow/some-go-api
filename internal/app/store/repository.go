package store

import "some-go-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
	LogAuthenticateAttempt(*model.AuthorizationLog) error
	FailedAttemptsCount(*model.User) (int, error)
	GetAuthorizeHistory(*model.User) ([]*model.AuthorizationLog, error)
	DeleteAuthorizeHistory(*model.User) error
	//FindByLoginPass(string, string) (*model.User, error)
}

type SessionRepository interface {
	Create(*model.Session) error
	Find(string) (*model.Session, error)
}

