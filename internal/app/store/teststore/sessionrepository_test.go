package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"some-go-api/internal/app/store/teststore"
	"testing"
	"time"
)

func TestSessionRepository_Create(t *testing.T) {
	s := teststore.New()

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	testSession := model.NewSession(testUser)
	s.Session().Create(testSession)
	assert.NoError(t, s.Session().Create(testSession))
}

func TestSessionRepository_FindByToken(t *testing.T) {
	s := teststore.New()

	_, err := s.Session().FindByToken("wrong_token")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	s.User().Create(testUser)
	testSession := model.NewSession(testUser)
	err = s.Session().Create(testSession)
	assert.NoError(t, err)

	foundSession, err := s.Session().FindByToken(testSession.Token)
	assert.NoError(t, err)
	assert.NotNil(t, foundSession)
	assert.WithinDuration(t, testSession.ExpirationTime, foundSession.ExpirationTime, time.Second*time.Duration(3))
}
