package teststore

import (
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
)

type SessionRepository struct {
	store *Store
	sessions map[int]*model.Session
}

func (r *SessionRepository) Create(s *model.Session) error {
	s.CreateToken()
	s.ID = len(r.sessions)
	r.sessions[s.ID] = s
	return nil
}

func (r *SessionRepository) FindByToken(token string) (*model.Session, error) {
	for _, s := range r.sessions {
		if s.Token == token {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
