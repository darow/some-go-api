package teststore

import (
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

//Store ...
type Store struct {
	userRepository *UserRepository
	sessionRepository *SessionRepository
}

//New ...
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
		authAttempts: make(map[int]*model.AuthorizationEvent),
	}
	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &SessionRepository{
		store: s,
		sessions: make(map[int]*model.Session),
	}
	return s.sessionRepository
}