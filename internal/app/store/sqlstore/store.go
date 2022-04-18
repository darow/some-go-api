package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"some-go-api/internal/app/store"
)

//Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	sessionRepository *SessionRepository
	authLogRepository *AuthLogRepository
}

//New Функция получения хранилища. У хранилища есть возможность взаимодействия с БД.
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &SessionRepository{
		store: s,
	}
	return s.sessionRepository
}


func (s *Store) AuthLog() store.AuthLogRepository {
	if s.authLogRepository != nil {
		return s.authLogRepository
	}

	s.authLogRepository = &AuthLogRepository{
		store: s,
	}
	return s.authLogRepository
}
