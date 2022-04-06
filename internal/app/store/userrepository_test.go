package store_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown("users")

	testUser := model.TestUser(t)
	_, err := s.User().FindByLogin(testUser.Login)
	assert.Error(t, err)
	s.User().Create(testUser)

	foundUser, err := s.User().FindByLogin(testUser.Login)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, testUser.Login == foundUser.Login, true)
}

func TestUserRepository_CheckPass(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown("users")

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	u, err := s.User().FindByLogin(testUser.Login)
	match, err := s.User().CheckPass(u, testUser.Password)

	assert.NoError(t, err)
	assert.True(t, match)
}

func TestUserRepository_FindByLoginPass(t *testing.T) {
	s, teardown := store.TestStore(t, psqlInfo)
	defer teardown("users")

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	foundedUser, err := s.User().FindByLoginPass(testUser.Login, testUser.Password)
	assert.NoError(t, err)
	assert.NotNil(t, foundedUser)
	assert.Equal(t, testUser.Login == foundedUser.Login, true)
}

