package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"some-go-api/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	assert.NoError(t, s.User().Create(model.TestUser(t)))
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	_, err := s.User().FindByLogin(testUser.Login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.User().Create(testUser)

	foundUser, err := s.User().FindByLogin(testUser.Login)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, testUser.Login == foundUser.Login, true)
}

func TestUserRepository_FindByLoginPass(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	foundedUser, err := s.User().FindByLoginPass(testUser.Login, testUser.Password)
	assert.NoError(t, err)
	assert.NotNil(t, foundedUser)
	assert.Equal(t, testUser.Login == foundedUser.Login, true)
}


