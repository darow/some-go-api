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

	testSession := model.TestSession(t)
	testSession.UserID = testUser.ID
	assert.NoError(t, s.Session().Create(testSession))
}

func TestSessionRepository_FindByToken(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, psqlInfo)
	defer teardown("sessions", "users")
	s := sqlstore.New(db)

	testSession := model.TestSession(t)
	_, err := s.Session().FindByToken(testSession.Token)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	s.User().Create(testUser)

	testSession.UserID = testUser.ID
	s.Session().Create(testSession)

	foundSession, err := s.Session().FindByToken(testSession.Token)
	assert.NoError(t, err)
	assert.NotNil(t, foundSession)
	assert.WithinDuration(t, testSession.ExpirationTime, foundSession.ExpirationTime, time.Second*time.Duration(3))
}