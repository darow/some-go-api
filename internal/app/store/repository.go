package store

import "some-go-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
	LogAuthenticateAttempt(*model.AuthorizationEvent) error
	FailedAttemptsCount(*model.User) (int, error)
	//FindByLoginPass(string, string) (*model.User, error)
}

type SessionRepository interface {
	Create(*model.Session) error
	FindByToken(string) (*model.Session, error)
}

