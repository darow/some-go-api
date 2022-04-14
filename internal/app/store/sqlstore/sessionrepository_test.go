package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"some-go-api/internal/app/model"
	"some-go-api/internal/app/store"
	"some-go-api/internal/app/store/sqlstore"
	"testing"
	"time"
)

func TestSessionRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("sessions", "users")
	s := sqlstore.New(db)

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	testSession := model.NewSession(testUser)
	assert.NoError(t, s.Session().Create(testSession))
}

func TestSessionRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("sessions", "users")
	s := sqlstore.New(db)

	_, err := s.Session().Find("wrong_token")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	testSession := model.NewSession(testUser)
	err = s.Session().Create(testSession)
	assert.NoError(t, err)

	foundSession, err := s.Session().Find(testSession.Token)
	assert.NoError(t, err)
	assert.NotNil(t, foundSession)
	assert.WithinDuration(t, testSession.ExpirationTime, foundSession.ExpirationTime, time.Second*time.Duration(3))
}