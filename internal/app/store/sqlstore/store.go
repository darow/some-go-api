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
}

//New ...
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
