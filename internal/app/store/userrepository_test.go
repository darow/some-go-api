package store_test

import (
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Login:             "test_login",
		EncryptedPassword: "asdf",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown()

	login := "test_login"
	_, err := s.User().FindByLogin(login)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Login:            	login,
		EncryptedPassword: "asdf",
	})
	u, err := s.User().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, login==u.Login, true)
}
