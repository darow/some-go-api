package store

import "some-go-api/internal/app/model"

//UserRepository Интерфейс для взаимодействия с пользователями в БД.
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
}

//SessionRepository Интерфейс для взаимодействия с сессиями в БД.
type SessionRepository interface {
	Create(*model.Session) error
	Find(string) (*model.Session, error)
}

//AuthLogRepository Интерфейс для взаимодействия с логами авторизации в БД.
type AuthLogRepository interface {
	LogAuthenticateAttempt(*model.AuthorizationLog) error
	FailedAttemptsCount(*model.User) (int, error)
	GetAuthorizeHistory(*model.User) ([]*model.AuthorizationLog, error)
	DeleteAuthorizeHistory(*model.User) error
}

